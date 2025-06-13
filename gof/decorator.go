package gof

import (
	"log"
)

type Handler1 func(s string)

// HandlerDecorator mapping Handler -> Handler
type HandlerDecorator func(Handler1) Handler1

var greetingHandler HandlerDecorator = func(handler Handler1) Handler1 {
	return func(s string) {
		s2 := "hello, " + s + ", welcome"
		log.Printf("[greetingHandler] s from %s -> %s\n", s, s2)
		handler(s2)
		log.Println("[greetingHandler]", "ending")
	}
}

var logHandler HandlerDecorator = func(handler Handler1) Handler1 {
	return func(s string) {
		handler(s)
		log.Println("[logHandler]", s)
	}
}

func CompositeHandler(handler Handler1, decorators ...HandlerDecorator) Handler1 {
	var h = handler
	for i := len(decorators) - 1; i >= 0; i-- {
		h = decorators[i](h)
	}
	return h
}
