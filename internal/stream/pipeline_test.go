package stream

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}
