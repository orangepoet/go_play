package main

import "fmt"

type Foo struct {
	val int
}

func (f *Foo) getVal() {
	fmt.Println("getVal->", f.val)
}

func main() {
	foos := make([]Foo, 0)
	change(&foos)

	for _, each := range foos {
		fmt.Println(each)
	}
}

func change(fs *[]Foo) {
	*fs = append(*fs, Foo{val: 10})
}
