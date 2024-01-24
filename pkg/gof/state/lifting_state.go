package state

import (
	"log"
	"time"
)

type LiftingState int

const (
	_ LiftingState = iota
	STATE_OPEN
	STATE_RUNNING
	STATE_STOP
)

func (s LiftingState) String() string {
	switch s {
	case STATE_OPEN:
		return "STATE_OPEN"
	case STATE_RUNNING:
		return "STATE_RUNNING"
	case STATE_STOP:
		return "STATE_STOP"
	default:
		return "UNKNOWN"
	}
}

var handlerMap map[LiftingState]CommandHandler

func init() {
	handlerMap = make(map[LiftingState]CommandHandler)
	handlerMap[STATE_OPEN] = handleWhenOpen
	handlerMap[STATE_STOP] = handleWhenStop
	handlerMap[STATE_RUNNING] = handleWhenRunning
}

// handleWhenRunning 运行时的处理
func handleWhenRunning(l *Lifting, c Command) {
	switch c.commandType {
	default:
		l.ignore(c)
	}
}

// handleWhenOpen 打开时的处理
func handleWhenOpen(l *Lifting, c Command) {
	switch c.commandType {
	case COMMAND_CLOSE:
		l.State = STATE_STOP
	default:
		l.ignore(c)
	}
}

// handleWhenStop 停止时的处理
func handleWhenStop(l *Lifting, c Command) {
	switch c.commandType {
	case COMMAND_OPEN:
		l.State = STATE_OPEN
	case COMMAND_RUN:
		l.State = STATE_RUNNING
		l.wg.Add(1)
		go func() {
			defer l.wg.Done()
			time.Sleep(5 * time.Second)
			l.State = STATE_STOP
			log.Printf("running terminal, now state: %s \n", l.State)
		}()
	default:
		l.ignore(c)
	}
}
