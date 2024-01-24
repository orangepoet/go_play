package observer

import (
	"fmt"
	"testing"
	"time"
)

type Observer1 struct {
}

func (obs *Observer1) listen(event Event) {
	fmt.Printf("ob1 Received event: %s\n", event.Message)
}

func TestEvent(t *testing.T) {

	// 创建事件发射器
	emitter := &EventEmitter{
		Handlers: make(map[string][]EventHandler),
	}

	// 注册事件处理函数
	ob2 := func(event Event) {
		fmt.Printf("ob2 Received event: %s\n", event.Message)
	}
	emitter.On("customEvent", ob2)
	emitter.On("customEvent", func(event Event) {
		fmt.Printf("ob3 Received event: %s\n", event.Message)
	})
	observer1 := &Observer1{}
	var handler3 EventHandler = observer1.listen
	emitter.On("customEvent", handler3)

	emitter.Off("customEvent", ob2)

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
