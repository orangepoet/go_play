package stream

import "go.uber.org/zap"

// Log A log.
var Log, _ = zap.NewProduction()

// Recover when panic happen.
func Recover() {
	if e := recover(); e != nil {
		Log.Error("stream failed", zap.Any("error", e))
	}
}

// NewGoroutine New a Safe goroutine.
func NewGoroutine(f func()) {
	defer Recover()
	f()
}
