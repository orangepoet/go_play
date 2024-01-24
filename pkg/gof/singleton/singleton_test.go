package singleton

import (
	"sync"
	"testing"
)

func TestObserver(t *testing.T) {
	var wg sync.WaitGroup

	ss := make([]*Singleton, 0)
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			inst := Instance()
			ss = append(ss, inst)
		}()
	}
	wg.Wait()

	var prev *Singleton
	for _, e := range ss {
		if prev == nil {
			prev = e
		} else if prev != e {
			t.Fatal("not singleton")
		}
	}
}
