package main

import "fmt"

type Function1 func(x, y int) int

func main() {
	var f Function1 = func(x, y int) int {
		return x + y
	}
	fmt.Printf("%v\n", f)
	fmt.Printf("%p\n", f)
	getFuncAddr(f)
}

func getFuncAddr(f Function1) {
	fmt.Printf("%v\n", f)
	fmt.Printf("%p\n", f)
}
