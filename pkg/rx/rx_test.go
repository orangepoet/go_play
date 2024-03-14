package rx

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"log"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	observable := rxgo.Just("Hello, World!")()
	//ch := observable.Observe()
	//item := <-ch
	//fmt.Println(item.V)

	observable.ForEach(func(v interface{}) {
		fmt.Printf("received: %v\n", v)
	}, func(err error) {
		fmt.Printf("error: %e\n", err)
	}, func() {
		fmt.Println("observable is closed")
	})
}

func TestMap(t *testing.T) {
	observe := rxgo.Just(1, 2, 3)().
		Map(func(_ context.Context, i interface{}) (interface{}, error) {
			return i.(int) * 10, nil
		}).Observe()
	for item := range observe {
		fmt.Println(item.V)
	}
}

func TestProducer(t *testing.T) {
	observable := rxgo.
		Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
			for i := 0; ; i++ {
				next <- rxgo.Of(i)
				time.Sleep(time.Second)
			}
		}}).
		Filter(func(i interface{}) bool {
			return i.(int)%2 == 0
		})

	subscriber := observable.
		ForEach(func(i interface{}) {
			log.Println(i.(int))
		}, func(_ error) {}, func() {})

	<-subscriber
}

func TestFromChannel(t *testing.T) {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()
	observable := rxgo.FromChannel(ch)
	for e := range observable.Observe() {
		log.Println(e.V)
	}
}

func TestInterval(t *testing.T) {
	observable := rxgo.Interval(rxgo.WithDuration(1 * time.Second))
	for item := range observable.Observe() {
		log.Println(item.V)
	}
}

func TestHot(t *testing.T) {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func TestCold(t *testing.T) {
	observable := rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func TestConnectable(t *testing.T) {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy())

	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("First observer: %d\n", i)
	})

	time.Sleep(3 * time.Second)
	fmt.Println("before subscribe second observer")

	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("Second observer: %d\n", i)
	})

	observable.Connect(context.Background())
	time.Sleep(3 * time.Second)
}
