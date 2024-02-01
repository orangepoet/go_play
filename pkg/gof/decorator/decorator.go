package decorator

import "fmt"

type Handler func(s string)

type HandlerDecorator func(Handler) Handler

var greetingHandler HandlerDecorator = func(handler Handler) Handler {
	return func(s string) {
		s = "hello, " + s + ", welcome"
		handler(s)
		fmt.Println("ending")
	}
}

func CompositeHandler(handler Handler, decorators ...HandlerDecorator) Handler {
	var h Handler
	for i := len(decorators) - 1; i >= 0; i-- {
		h = decorators[i](handler)
	}
	return h
}
