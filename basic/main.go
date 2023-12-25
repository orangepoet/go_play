package main

import "fmt"

func main() {
	map1 := make(map[string][]int)
	map1["a"] = []int{1, 2, 3}
	v := map1["b"]
	for _, e := range v {
		fmt.Println(e)
	}
}
