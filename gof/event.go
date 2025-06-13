package gof

import (
	"fmt"
)

// Event 类型表示一个事件
type Event struct {
	Message string
}

// EventHandler 事件处理器
type EventHandler func(event Event)

// EventEmitter 类型表示事件发射器
type EventEmitter struct {
	Handlers map[string][]EventHandler
}

// On 方法用于注册事件处理函数
func (emitter *EventEmitter) On(eventType string, handler EventHandler) {
	if emitter.Handlers == nil {
		emitter.Handlers = make(map[string][]EventHandler)
	}

	emitter.Handlers[eventType] = append(emitter.Handlers[eventType], handler)
}

func (emitter *EventEmitter) Off(eventType string, handler EventHandler) {
	handlers := emitter.Handlers[eventType]
	for i, e := range handlers {
		addr0 := fmt.Sprintf("%p", e)
		addr1 := fmt.Sprintf("%p", handler)
		if addr0 == addr1 {
			emitter.Handlers[eventType] = append(handlers[:i], handlers[i+1:]...)
			break
		}
	}
}

// Emit 方法用于发射事件
func (emitter *EventEmitter) Emit(eventType string, event Event) {
	handlers, ok := emitter.Handlers[eventType]
	if ok {
		for _, handler := range handlers {
			handler(event)
		}
	}
}
