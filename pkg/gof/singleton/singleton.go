package singleton

import "sync"

// Singleton 单例
type Singleton struct {
}

var instance *Singleton
var once sync.Once

// Instance 单例
func Instance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
