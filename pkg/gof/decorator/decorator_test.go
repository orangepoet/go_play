package decorator

import (
	"fmt"
	"testing"
)

func Test_decorator(t *testing.T) {
	var handler = func(s string) {
		fmt.Println(s)
	}
	decorators := make([]HandlerDecorator, 0)
	decorators = append(decorators, greetingHandler)
	ch := CompositeHandler(handler, decorators...)
	ch("orange_cheng")
}
