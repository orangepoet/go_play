package gof

import (
	"fmt"
	"sync"
)

// Observer 接口定义了观察者的方法
type Observer interface {
	Update(message string)
}

// Subject 定义了主题的结构，包括观察者列表和互斥锁
type Subject struct {
	observers []Observer
	mutex     sync.Mutex
}

// RegisterObserver 用于注册观察者
func (s *Subject) RegisterObserver(observer Observer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.observers = append(s.observers, observer)
}

// NotifyObservers 用于通知所有注册的观察者
func (s *Subject) NotifyObservers(message string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

// ConcreteObserver 实现了 Observer 接口的具体观察者
type ConcreteObserver struct {
	name string
}

// Update 实现了观察者的 Update 方法
func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("[%s] Received message: %s\n", o.name, message)
}
