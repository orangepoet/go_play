package oop

import (
	"fmt"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {
	// 创建事件发射器
	emitter := &EventEmitter{}

	// 注册事件处理函数
	emitter.On("customEvent", func(event Event) {
		fmt.Printf("handler1 Received event: %s\n", event.Message)
	})
	emitter.On("customEvent", func(event Event) {
		fmt.Printf("handler2 Received event: %s\n", event.Message)
	})

	// 模拟发射事件
	go func() {
		for i := 1; i <= 3; i++ {
			event := Event{Message: fmt.Sprintf("Event %d occurred", i)}
			emitter.Emit("customEvent", event)
			time.Sleep(time.Second)
		}
	}()

	// 主程序休眠一段时间，以等待事件处理
	time.Sleep(5 * time.Second)
}
