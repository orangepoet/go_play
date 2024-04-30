package gof

import "testing"

func TestHandler1_Handle(t *testing.T) {
	handler1 := ConcreteHandler{HandlerNode{name: "handler1"}}
	handler2 := ConcreteHandler2{HandlerNode{name: "handler2"}}
	handler3 := ConcreteHandler3{HandlerNode{name: "handler3"}}
	handler1.next = &handler2
	handler2.next = &handler3

	handler1.Handle(Event3{i: 1, s: "a"})
}
