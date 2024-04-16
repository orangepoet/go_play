package fp

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func Chan[T any](xs []T) <-chan T {
	out := make(chan T)
	go func() {
		for _, n := range xs {
			out <- n
			//log.Println("Chan <-", n)
			//time.Sleep(5 * time.Millisecond)
		}
		close(out)
	}()
	return out
}

type PipeFunc[T any] func(<-chan T) <-chan T

func Pipeline[T any](xs []T, pipeFns ...PipeFunc[T]) <-chan T {
	ch := Chan(xs)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}

func Map0[T, U any](mapTo func(x T) U) func(in <-chan T) <-chan U {
	return func(in <-chan T) <-chan U {
		out := make(chan U)
		go func() {
			for x := range in {
				out <- mapTo(x)
				fmt.Println("Map0", x)
			}
			close(out)
		}()
		return out
	}
}

func Merge[T any](ins []<-chan T) <-chan T {
	out := make(chan T)
	for _, in := range ins {
		go func(c <-chan T) {
			for e := range c {
				out <- e
			}
		}(in)
	}
	return out
}

func Merge1[T any](c1, c2 <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		for {
			out <- <-c1
		}
	}()
	go func() {
		for {
			out <- <-c2
		}
	}()

	return out
}

func Merge0[T any](cs ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan T) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func FlatMap[T, U any](flatMap func(x T) []U) func(in <-chan T) <-chan U {
	return func(in <-chan T) <-chan U {
		out := make(chan U)
		time.Sleep(3 * time.Second)
		go func() {

			for x := range in {
				log.Println("FlatMap <-", x)
				for _, u := range flatMap(x) {
					log.Println("FlatMap.loop <-", u)
					out <- u
					time.Sleep(20 * time.Millisecond)
				}
			}
			close(out)
		}()
		return out
	}
}

func Reduce[T any](reduce func(a, b T) T) func(in <-chan T) T {
	return func(in <-chan T) T {
		var wg sync.WaitGroup
		wg.Add(1)
		var c T
		go func() {
			for e1 := range in {
				c = reduce(c, e1)
			}
			time.Sleep(time.Second)
			wg.Done()
		}()
		wg.Wait()
		return c
	}
}

// Foreach Sync Execute method
func Foreach[T any](handle func(x T)) func(in <-chan T) {
	return func(in <-chan T) {
		for x := range in {
			fmt.Println("handle", x)
			handle(x)
		}
	}
}

func Filter0[T any](predicate func(x T) bool) func(in <-chan T) <-chan T {
	return func(in <-chan T) <-chan T {
		out := make(chan T)
		go func() {
			for x := range in {
				if predicate(x) {
					out <- x
				}
			}
			close(out)
		}()
		return out
	}
}

// Split 集合分割
func Split[T any](predicate func(x T) bool) func(xs []T) ([]T, []T) {
	return func(xs []T) ([]T, []T) {
		left := make([]T, 0)
		right := make([]T, 0)
		if len(xs) == 0 {
			return left, right
		}

		for _, x := range xs {
			if predicate(x) {
				left = append(left, x)
			} else {
				right = append(right, x)
			}
		}
		return left, right
	}

}

// AnyMatch 匹配
func AnyMatch[T any](predicate func(x T) bool) func(xs []T) bool {
	return func(xs []T) bool {
		for _, x := range xs {
			if predicate(x) {
				return true
			}
		}
		return false
	}
}

// Filter 过滤
func Filter[T any](predicate func(x T) bool) func(xs []T) []T {
	return func(xs []T) []T {
		result := make([]T, 0)
		if len(xs) == 0 {
			return result
		}
		for _, x := range xs {
			if predicate(x) {
				result = append(result, x)
			}
		}
		return result
	}
}

// FindFirst 找到首个满足条件
func FindFirst[T any](predicate func(x T) bool) func(xs []T) (T, bool) {
	return func(xs []T) (T, bool) {
		for _, x := range xs {
			if predicate(x) {
				return x, true
			}
		}
		var x T
		return x, false
	}
}

// Map 集合映射
func Map[T, U any](mapTo func(x T) U) func(xs []T) []U {
	return func(ts []T) []U {
		us := make([]U, 0)
		for _, x := range ts {
			us = append(us, mapTo(x))
		}
		return us
	}
}

// GroupBy 分组
func GroupBy[T any, K comparable](keyFunc func(x T) K) func(xs []T) map[K][]T {
	return func(xs []T) map[K][]T {
		group := make(map[K][]T, 0)
		for _, x := range xs {
			group[keyFunc(x)] = append(group[keyFunc(x)], x)
		}
		return group
	}
}
