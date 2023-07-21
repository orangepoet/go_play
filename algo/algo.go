// algorithm impl by go
package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

// jump game
func canJump(nums []int) bool {
	m := 0
	last := len(nums)
	for idx, e := range nums {
		if idx > m {
			return false
		}
		next := idx + e
		if next >= last {
			return true
		}
		m = max(m, next)
	}
	return m > len(nums)
}

func deleteDuplicates(head *ListNode) *ListNode {
	return head
}

func makeList(arr []int) *ListNode {
	var next ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		cur := ListNode{arr[i], &next}
		next = cur
	}
	return &next
}

func (head *ListNode) toArr() []int {
	r := []int{}
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	return r
}

func equals(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
