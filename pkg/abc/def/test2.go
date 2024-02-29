package def

import "go_play/pkg/abc"

type Woo struct {
	foo *abc.Foo
}

func NewWoo(foo *abc.Foo) *Woo {
	return &Woo{foo: foo}
}
