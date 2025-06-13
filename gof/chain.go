package gof

import "fmt"

type Event0 interface {
}

type SomeEvent0 struct {
}

type Handler interface {
	Handle(e Event0)
}

type HandlerNode struct {
	name string
	next Handler
}

func (n *HandlerNode) fireNext(e Event0) {
	if n.next != nil {
		fmt.Println("fire next")
		n.next.Handle(e)
	} else {
		fmt.Println("no next left")
	}
}

type Event1 struct {
	s string
}

type ConcreteHandler struct {
	HandlerNode
}

func (c *ConcreteHandler) Handle(e Event0) {
	// if can handle, handle it and not fire next.
	switch e.(type) {
	case Event1, Event2:
		fmt.Printf("%s handle event: %v\n", c.name, e)
	}
	c.fireNext(e)
}

type Event2 struct {
	val int
}

type ConcreteHandler2 struct {
	HandlerNode
}

func (c *ConcreteHandler2) Handle(e Event0) {
	if e2, ok := e.(Event2); ok {
		fmt.Printf("%s handle event: %v\n", c.name, e2)
	} else {
		c.fireNext(e)
	}
}

type Event3 struct {
	s string
	i int
}

type ConcreteHandler3 struct {
	HandlerNode
}

func (c *ConcreteHandler3) Handle(e Event0) {
	if e3, ok := e.(Event3); ok {
		fmt.Printf("%s handle event: %v\n", c.name, e3)
	}
	c.fireNext(e)
}

func init() {

}
