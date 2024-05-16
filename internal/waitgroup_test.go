package internal

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	var wg sync.WaitGroup

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			execTime := r.Intn(10)
			fmt.Println("Goroutine", id, "started", "execTime", execTime)
			time.Sleep(time.Duration(execTime) * time.Second)
			fmt.Println("Goroutine", id, "completed")
		}(i)
	}

	wg.Wait()
	fmt.Println("All Goroutines completed")
}
