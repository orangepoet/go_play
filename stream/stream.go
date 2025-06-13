package stream

import (
	"sync"
)

type (
	// FilterFunc defines the method to filter a Stream.
	FilterFunc func(item interface{}) bool
	// MapFunc defines the method to map each element to another object in a Stream.
	MapFunc func(item interface{}) interface{}
	// ForEachFunc defines the method to handle each element in a Stream.
	ForEachFunc func(item interface{})
	// WalkFunc defines the method to walk through all the elements in a Stream.
	WalkFunc func(item interface{}, pipe chan<- interface{})
)

type Stream struct {
	source <-chan interface{}
}

func Range(source <-chan interface{}) *Stream {
	return &Stream{
		source: source,
	}
}

func Of(items ...interface{}) *Stream {
	source := make(chan interface{}, len(items))
	go func() {
		for _, item := range items {
			source <- item
		}
		close(source)
	}()
	return Range(source)
}

func (s *Stream) Walk(f WalkFunc, opts ...Option) *Stream {
	option := loadOptions(opts...)
	pipe := make(chan interface{}, option.workSize)
	go func() {
		var wg sync.WaitGroup
		pool := make(chan struct{}, option.workSize)

		for {
			pool <- struct{}{}
			item, ok := <-s.source
			if !ok {
				<-pool
				break
			}

			wg.Add(1)
			// better to safely run caller defined method
			go NewGoroutine(func() {
				defer func() {
					wg.Done()
					<-pool
				}()
				f(item, pipe)
			})
		}
		wg.Wait()
		close(pipe)
	}()

	return Range(pipe)
}

func loadOptions(options ...Option) *Options {
	op := new(Options)
	for _, option := range options {
		option(op)
	}
	// set the default pool size
	if op.workSize <= 0 {
		op.workSize = 1
	}
	return op
}

func (s *Stream) Map(fn MapFunc, opts ...Option) *Stream {
	return s.Walk(func(item interface{}, pipe chan<- interface{}) {
		pipe <- fn(item)
	}, opts...)
}

func (s *Stream) Filter(fn FilterFunc, opts ...Option) *Stream {
	return s.Walk(func(item interface{}, pipe chan<- interface{}) {
		if fn(item) {
			pipe <- item
		}
	}, opts...)
}

// AllMatch Returns whether all elements of this stream match the provided predicate.
// May not evaluate the predicate on all elements if not necessary for determining the result.
// If the stream is empty then true is returned and the predicate is not evaluated.
func (s *Stream) AllMatch(f func(item interface{}) bool) (isFind bool) {
	isFind = true
	for item := range s.source {
		if !f(item) {
			isFind = false
			return
		}
	}
	return
}

func (s *Stream) Foreach(f ForEachFunc) {
	items := make([]interface{}, 0)
	for item := range s.source {
		f(item)
		items = append(items, item)
	}

}

func (s *Stream) Collect() []interface{} {
	c := make([]interface{}, 0)
	for n := range s.source {
		c = append(c, n)
	}
	return c
}

// Chan Returns a channel of Stream.
func (s *Stream) Chan() <-chan interface{} {
	return s.source
}
