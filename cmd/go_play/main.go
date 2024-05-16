package main

import "fmt"

func main() {
	var foo *Foo

	foo = nil

	fmt.Printf("%v", foo)
}

type Foo struct {
}
