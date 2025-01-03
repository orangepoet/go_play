package internal

import (
	"math/rand"
	"strings"
)

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

func min(a, b int) int {
	if a < b {
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

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func preOrder(tn *TreeNode) []int {
	if tn == nil {
		return nil
	}
	ans := []int{tn.val}
	ans = append(ans, preOrder(tn.left)...)
	ans = append(ans, preOrder(tn.right)...)

	return ans
}

func addList(left, right *ListNode) *ListNode {
	dummp := &ListNode{}
	p := dummp

	proceed := false
	for left != nil || right != nil {
		sum := 0
		if left != nil {
			sum += left.Val
			left = left.Next
		}
		if right != nil {
			sum += right.Val
			right = right.Next
		}
		if proceed {
			sum++
		}
		if sum >= 10 {
			proceed = true
			sum %= 10
		} else {
			proceed = false
		}
		p.Next = &ListNode{Val: sum}
		p = p.Next
	}
	if proceed {
		p.Next = &ListNode{Val: 1}
	}
	return dummp.Next
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	generateParenthesis0(&ans, n, n, "")
	return ans
}

func generateParenthesis0(ans *[]string, l, r int, s string) {
	if l == 0 && r == 0 {
		*ans = append(*ans, s)
		return
	}
	if l > 0 {
		generateParenthesis0(ans, l-1, r, s+"(")
	}
	if r > l {
		generateParenthesis0(ans, l, r-1, s+")")
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	cur, next := dummy, head

	for i := 0; i < n; i++ {
		if next == nil {
			break
		}
		cur = next
		next = next.Next
	}
	cur.Next = next.Next
	return dummy.Next
}

func makeChange(n int, coins []int) int {
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		changes := n + 1
		for _, coin := range coins {
			if coin <= i && dp[i-coin]+1 < changes {
				changes = dp[i-coin] + 1
			}
		}
		dp[i] = changes
	}
	if dp[n] > n {
		return -1
	}
	return dp[n]
}

func patternMatch(pattern, words string) bool {
	// "abba"  | "dog cat cat dog"
	mapping := make(map[string]uint8)
	stored := make(map[uint8]bool)
	wordSlice := strings.Split(words, " ")
	for idx, w := range wordSlice {
		if value, ok := mapping[w]; ok {
			if value != pattern[idx] {
				return false
			}
		} else {
			if ok := stored[pattern[idx]]; ok {
				return false
			}
			mapping[w] = pattern[idx]
			stored[pattern[idx]] = true
		}
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.left), maxDepth(root.right))
}

func isMirror(node *TreeNode) bool {
	return isMirror0(node, node)
}

func isMirror0(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.val == r.val && isMirror0(l.left, r.right) && isMirror0(l.right, r.left)
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

func searchRange(nums []int, target int) []int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			start, end := mid, mid
			for start-1 >= 0 && nums[start-1] == target {
				start--
			}
			for end+1 <= len(nums)-1 && nums[end+1] == target {
				end++
			}
			return []int{start, end}
		}
	}
	return []int{-1, -1}
}

// public List<Integer> rightSideView(TreeNode root) {
func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)

	if root == nil {
		return ret
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	start, end := 0, 0
	for start <= end {
		ret = append(ret, queue[end].val)
		for i := start; i <= end; i++ {
			if queue[i].left != nil {
				queue = append(queue, queue[i].left)
			}
			if queue[i].right != nil {
				queue = append(queue, queue[i].right)
			}
		}
		start, end = end+1, len(queue)-1
	}

	return ret
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 0
	m := make(map[byte]int)
	for start, end := 0, 0; end < len(s); end++ {
		x := s[end]
		if v, ok := m[x]; ok {
			start = v + 1
		} else {
			width := end - start + 1
			if width > max {
				max = width
			}
		}
		m[x] = end
	}

	return max
}

func numIslands(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	y, x := len(grid), len(grid[0])
	count := 0
	for h := 0; h < x; h++ {
		for v := 0; v < y; v++ {
			if grid[v][h] == 1 {
				numIslands0(grid, h, v, x, y)
				count++
			}
		}
	}
	return count
}

func numIslands0(grid [][]int, x, y, xl, yl int) {
	if x >= xl || y >= yl {
		return
	}
	if grid[y][x] != 1 {
		return
	}
	grid[y][x] = 0
	numIslands0(grid, x+1, y, xl, yl)
	numIslands0(grid, x, y+1, xl, yl)
}

// WeightSelect 权重选取
func weightSelect(vw map[int]int) (int, bool) {
	total := 0
	for _, w := range vw {
		total += w
	}
	target := rand.Intn(total)
	cur := 0
	for v, w := range vw {
		cur += w
		if cur > target {
			return v, true
		}
	}
	return 0, false
}
