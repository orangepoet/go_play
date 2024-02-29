package fp

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAnyMatch(t *testing.T) {
	s := []int{1, 2, 3, 4}
	b := AnyMatch(func(x int) bool {
		return x > 5
	})(s)
	fmt.Println(b)
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4}
	filtered := Filter(func(x int) bool {
		return x%2 == 0
	})(s)
	fmt.Println(filtered)
}

func TestFindFirst(t *testing.T) {
	s := []int{1, 2, 3, 4}
	ele, find := FindFirst(func(x int) bool {
		return x == 1
	})(s)
	fmt.Println(ele, find)
}

func TestGroupBy(t *testing.T) {
	s := []int{1, 2, 3, 4}
	m := GroupBy(func(x int) bool {
		return x%2 == 0
	})(s)
	fmt.Println(m)
}

func TestMap(t *testing.T) {
	s := []int{1, 2, 3, 4}
	strings := Map(func(x int) string {
		return "e:" + strconv.Itoa(x)
	})(s)
	fmt.Println(strings)
}

func TestSplit(t *testing.T) {
	s := []int{1, 2, 3, 4}
	left, right := Split(func(x int) bool {
		return x%2 == 0
	})(s)
	fmt.Println(left, right)
}
