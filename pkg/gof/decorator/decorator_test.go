package decorator

import (
	"log"
	"testing"
)

func Test_decorator(t *testing.T) {
	var handler = func(s string) {
		log.Println("[main]", s)
	}
	decorators := make([]HandlerDecorator, 0)
	decorators = append(decorators, greetingHandler, logHandler)
	ch := CompositeHandler(handler, decorators...)
	ch("orange_cheng")
}
