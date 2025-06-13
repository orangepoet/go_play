package play

import (
	"fmt"
	"sort"
)

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) Copy() Sequence {
	c := make(Sequence, 0, len(s))
	return append(c, s...)
}

func (s Sequence) String() string {
	s = s.Copy()
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}

type IService interface {
	DoSth1()
	DoSth2()
	mustImplementIService()
}

type MyService struct {
}

func (receiver MyService) DoSth1() {
	fmt.Println("MyService.DoSth1")
}

func (receiver MyService) DoSth2() {
	fmt.Println("MyService.DoSth2")
}

func (receiver MyService) mustImplementIService() {

}

type ExtendedService struct {
	MyService
}

func (receiver ExtendedService) DoSth1() {
	fmt.Println("ExtendedService.DoSth1")
}
