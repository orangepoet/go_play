package observer

import (
	"fmt"
	"testing"
	"time"
)

func TestObserver(t *testing.T) {
	// 创建主题
	subject := &Subject{}

	// 创建观察者并注册到主题
	observer1 := &ConcreteObserver{name: "Observer1"}
	observer2 := &ConcreteObserver{name: "Observer2"}
	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	// 模拟事件通知
	go func() {
		for i := 1; i <= 3; i++ {
			message := fmt.Sprintf("Event %d occurred", i)
			subject.NotifyObservers(message)
			time.Sleep(time.Second)
		}
	}()

	// 主程序休眠一段时间，以便观察者有足够的时间接收通知
	time.Sleep(5 * time.Second)
}
