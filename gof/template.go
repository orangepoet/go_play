package gof

import "fmt"

type MyInter interface {
	step1()
	step2()
}

type Impl1 struct {
}

func templateMethod(inter MyInter) {
	fmt.Println("template method.. define main flow")
	inter.step1()
	inter.step2()
	fmt.Println("template method finish")
}

func (i *Impl1) step1() {
	fmt.Println("step1")
}

func (i *Impl1) step2() {
	fmt.Println("step2")
}
