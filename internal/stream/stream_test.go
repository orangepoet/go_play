package stream

import (
	"log"
	"testing"
)

func TestMap1(t *testing.T) {

	stream := Of(0, 1, 2, 3, 4).
		Map(func(item interface{}) interface{} {
			return item.(int) + 1
		}).
		Filter(func(item interface{}) bool {
			return item.(int) > 3
		})

	allMach := stream.AllMatch(func(item interface{}) bool {
		return item.(int)%2 == 0
	})
	log.Printf("allMatch: %t\n", allMach)

	// 拿不到元素了
	for n := range stream.Chan() {
		log.Printf(" <- %v", n)
	}

	collect := Of(0, 1, 2, 3, 4).
		Map(func(item interface{}) interface{} {
			return item.(int) + 1
		}).
		Filter(func(item interface{}) bool {
			return item.(int) > 3
		}).Collect()
	allMatch2 := Of(collect...).AllMatch(func(item interface{}) bool {
		return item.(int)%2 == 0
	})
	log.Printf("allMatch2: %t\n", allMatch2)

	Of(collect...).Foreach(func(item interface{}) {
		log.Printf("foreach item: %d", item.(int))
	})
}
