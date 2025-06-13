package gof

import (
	"log"
	"sync"
)

type CommandType int

const (
	_ CommandType = iota
	COMMAND_RUN
	COMMAND_OPEN
	COMMAND_CLOSE
)

func (c CommandType) String() string {
	switch c {
	case COMMAND_CLOSE:
		return "COMMAND_CLOSE"
	case COMMAND_RUN:
		return "COMMAND_RUN"
	case COMMAND_OPEN:
		return "COMMAND_OPEN"
	default:
		return "UNKNOWN"
	}
}

type CommandHandler func(l *Lifting, c Command)

type Lifting struct {
	State LiftingState
	wg    sync.WaitGroup
}

type Command struct {
	commandType CommandType
	expectState LiftingState
}

func (c *Command) String() string {
	return "{ commandType: " + c.commandType.String() + "}"
}

func (l *Lifting) receive(c Command) {
	log.Printf("receive commmand: %s\n", c.commandType)
	handler := handlerMap[l.State]
	if handler == nil {
		log.Fatalf("no handle for state: %v", l.State)
		return
	}
	handler(l, c)
	log.Printf("lifting state: %v\n", l.State)
}

func (l *Lifting) WaitForStop() {
	l.wg.Wait()
}

func (l *Lifting) ignore(c Command) {
	log.Printf("lifting[%v] ignore command: %s\n", l.State, &c)
}
