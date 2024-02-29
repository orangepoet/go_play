package abc

import "go_play/pkg/abc/def"

type Foo struct {
	woo *def.Woo
}

func NewFoo() *Foo {
	return &Foo{}
}
