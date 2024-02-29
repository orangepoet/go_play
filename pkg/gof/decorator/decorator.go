package decorator

import (
	"log"
)

type Handler func(s string)

// HandlerDecorator mapping Handler -> Handler
type HandlerDecorator func(Handler) Handler

var greetingHandler HandlerDecorator = func(handler Handler) Handler {
	return func(s string) {
		s2 := "hello, " + s + ", welcome"
		log.Printf("[greetingHandler] s from %s -> %s\n", s, s2)
		handler(s2)
		log.Println("[greetingHandler]", "ending")
	}
}

var logHandler HandlerDecorator = func(handler Handler) Handler {
	return func(s string) {
		handler(s)
		log.Println("[logHandler]", s)
	}
}

func CompositeHandler(handler Handler, decorators ...HandlerDecorator) Handler {
	var h = handler
	for i := len(decorators) - 1; i >= 0; i-- {
		h = decorators[i](h)
	}
	return h
}
