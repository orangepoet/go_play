package internal

import (
	"fmt"
	"strconv"
	"testing"
	"time"
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

func TestMap0(t *testing.T) {
	s := []int{1, 2, 3, 4}
	//ans := Map(func(x int) int { return x * 2 })(Filter(func(x int) bool { return x%2 == 0 })(s))

	strings := Map(func(x int) string {
		return "e:" + strconv.Itoa(x)
	})(s)
	fmt.Println(strings)

	ans2 := Map(func(x int) int {
		return x + 1
	})(s)
	fmt.Println(ans2)
}

func TestSplit(t *testing.T) {
	s := []int{1, 2, 3, 4}
	left, right := Split(func(x int) bool {
		return x%2 == 0
	})(s)
	fmt.Println(left, right)
}

func TestStream(t *testing.T) {
	s := []int{1, 2, 3, 4}
	ret := Map0(func(x int) int { return x + 1 })(Chan(s))
	fmt.Println(ret)
}

func TestStreamPipeline(t *testing.T) {
	list := []int{1, 2, 3}
	sqFn := func(x int) int { return x * x }
	isOddFn := func(x int) bool {
		return x%2 == 0
	}
	ret2 := Pipeline(list, Map0(sqFn), Filter0(isOddFn))
	fmt.Println(ret2)
	//time.Sleep(100 * time.Second)
	for n := range ret2 {
		fmt.Println(n)
	}
}

func TestFlatMap(t *testing.T) {
	provinces := []Province{
		{name: "jiangxi", cities: []City{{"jiujiang"}, {"nanchang"}, {"jingdezhen"}}},
		{name: "zhejiang", cities: []City{{"hangzhou"}, {"wenzhou"}, {"ningbo"}}},
	}
	mapToCitys := func(p Province) []City {
		return p.cities
	}
	cities := Map0(func(c City) string { return c.name })(FlatMap(mapToCitys)(Chan(provinces)))
	//for n := range cities {
	//	fmt.Println("out", n)
	//}
	reduce := Reduce(func(acc, a string) string {
		if acc == "" {
			return a
		}
		return acc + "," + a
	})(cities)
	fmt.Println("reduce->", reduce)
	time.Sleep(60 * time.Second)
}

func TestFanInFanOut(t *testing.T) {
	in := Chan([]int{2, 3})
	addOneFn := func(x int) int { return x + 1 }
	c1 := Map0(addOneFn)(in)
	c2 := Map0(addOneFn)(in)
	for n := range Merge1(c1, c2) {
		fmt.Println(n)
	}
}

type City struct {
	name string
}

type Province struct {
	name   string
	cities []City
}

func TestCh(t *testing.T) {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	close(c)
	go func() {
		for {
			select {
			case in, ok := <-c:
				if !ok {
					return
				}
				fmt.Println(in)
			}
		}
	}()
	time.Sleep(10 * time.Second)
}
