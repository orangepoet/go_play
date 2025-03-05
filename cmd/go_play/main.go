package main

import (
	"errors"
	"fmt"

	"github.com/samber/lo"
)

type Foo struct {
	Name string
	Age  int
}

func main() {
	errs := make([]error, 0)
	for _, v := range lo.Range(10) {
		if e := makeFunc(v); e != nil {
			errs = append(errs, e)
		}
	}
	fe := errors.Join(errs...)
	fmt.Println(fe)


	for _, v:= range lo.Range(10) {
		fmt.Println(v)
	}



















}

func makeFunc(v int) error {
	if v%2 == 0 {
		return fmt.Errorf("even number: %d", v)
	}
	return nil

}

func Xx() string {
	return ""
}
