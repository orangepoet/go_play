package oop

// Event 类型表示一个事件
type Event struct {
	Message string
}

type EventHandler func(event Event)

// EventHandler 类型是处理事件的函数类型
//type EventHandler func(event Event)

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

// Emit 方法用于发射事件
func (emitter *EventEmitter) Emit(eventType string, event Event) {
	handlers, ok := emitter.Handlers[eventType]
	if ok {
		for _, handler := range handlers {
			handler(event)
		}
	}
}
