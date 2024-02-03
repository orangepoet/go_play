package algo

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// canJump: jump game
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
	var next *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		next = &ListNode{arr[i], next}
	}
	return next
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

func resolvePrime(n int) []int {
	ans := make([]int, 0)

	for left, f := n, 2; left > 2 && f <= left; {
		if left%f == 0 {
			ans = append(ans, f)
			left = left / f
		} else {
			f++
		}
	}
	return ans
}

func comb(n, m, k int) [][]int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	ans := make([][]int, 0)
	dfs(arr, m, k, 0, make([]int, 0), 0, &ans)
	return ans
}

func dfs(arr []int, left int, k int, sum int, c []int, next int, ans *[][]int) {
	if left == 0 && k == sum {
		*ans = append(*ans, c)
		return
	}
	for i := next; i < len(arr); i++ {
		dfs(arr, left-1, k, sum+arr[i], append(c, arr[i]), i+1, ans)
	}
}

func isValid(s string) bool {
	left, right := '{', '}'
	stack := make([]rune, 0)
	for _, c := range s {
		if c == left {
			stack = append(stack, c)
		} else if c == right {
			top := len(stack) - 1
			if top < 0 {
				return false
			}
			if stack[top] == left {
				stack = stack[:top]
			}
		}
	}
	return len(stack) == 0
}

func mergeListNode(list []*ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy

	for {
		selected := 0
		for i := range list {
			if list[i] == nil {
				continue
			}
			if list[selected] == nil || list[i].Val < list[selected].Val {
				selected = i
			}
		}
		if list[selected] == nil {
			break
		}
		prev.Next = list[selected]
		prev = list[selected]
		list[selected] = list[selected].Next
	}
	return dummy.Next
}
