package algo

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	t1 := canJump([]int{2, 3, 1, 1, 4})
	if !t1 {
		t.Errorf("expected true but false")
	}

	t2 := canJump([]int{3, 2, 1, 0, 4})
	if t2 {
		t.Errorf("expected false but true")
	}

	t3 := canJump([]int{1, 2, 3})
	if !t3 {
		t.Errorf("expected true but false")
	}
}

func TestDeleteDuplicates(t *testing.T) {
	r := deleteDuplicates(makeList([]int{1, 2, 3, 3, 4, 4, 5}))
	if !equals(r.toArr(), []int{1, 2, 5}) {
		t.Errorf("expected {1,2,5} but %d", r.toArr())
	}
}
