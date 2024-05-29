package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	bgCtx := context.Background()
	ctx, cancelFunc := context.WithTimeout(bgCtx, time.Second*2)
	defer func() {
		cancelFunc()
	}()

	ch := tryGetResult()
	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	case t := <-ch:
		fmt.Println(t)
		fmt.Println("received")
	}

	fmt.Println("done")
}

func tryGetResult() <-chan struct{} {
	c := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(c)
		fmt.Println("emit")
	}()
	return c
}
