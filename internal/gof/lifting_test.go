package gof

import (
	"log"
	"testing"
	"time"
)

func TestLifting(t *testing.T) {
	commands := []Command{
		{commandType: COMMAND_RUN, expectState: STATE_RUNNING},
		{commandType: COMMAND_OPEN, expectState: STATE_OPEN},
		{commandType: COMMAND_CLOSE, expectState: STATE_STOP},
		{commandType: COMMAND_RUN, expectState: STATE_RUNNING},
	}
	l := &Lifting{State: STATE_STOP}
	for _, each := range commands {
		l.receive(each)
		for l.State == STATE_RUNNING {
			log.Println("waiting for lifting to stop")
			time.Sleep(1 * time.Second)
		}
	}
	l.WaitForStop()
}
