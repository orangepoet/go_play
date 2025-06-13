package play

import (
	"math/rand"
	"strings"
)

// ListNode 定义链表节点结构
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 指向下一个节点的指针
}

// canJump 判断是否能跳到最后一个位置
// nums: 每个位置可以跳跃的最大长度
func canJump(nums []int) bool {
	m := 0 // 当前能到达的最远距离
	last := len(nums)
	for idx, e := range nums {
		if idx > m { // 如果当前位置超过了能到达的最远距离，返回false
			return false
		}
		next := idx + e
		if next >= last { // 如果能跳到或超过终点，返回true
			return true
		}
		m = max(m, next) // 更新最远可达距离
	}
	return m > len(nums)
}

// deleteDuplicates 删除链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	return head
}

// makeList 通过数组创建链表
func makeList(arr []int) *ListNode {
	var next *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		next = &ListNode{arr[i], next}
	}
	return next
}

// toArr 将链表转换为数组
func (head *ListNode) toArr() []int {
	r := []int{}
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	return r
}

// equals 判断两个整数数组是否相等
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

// max 返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// resolvePrime 分解质因数
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

// comb 计算组合数
// n: 总数, m: 选择数量, k: 目标和
func comb(n, m, k int) [][]int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	ans := make([][]int, 0)
	dfs(arr, m, k, 0, make([]int, 0), 0, &ans)
	return ans
}

// dfs 深度优先搜索辅助函数
func dfs(arr []int, left int, k int, sum int, c []int, next int, ans *[][]int) {
	if left == 0 && k == sum {
		*ans = append(*ans, c)
		return
	}
	for i := next; i < len(arr); i++ {
		dfs(arr, left-1, k, sum+arr[i], append(c, arr[i]), i+1, ans)
	}
}

// isValid 判断括号是否有效
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

// mergeListNode 合并多个有序链表
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

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// preOrder 前序遍历二叉树
func preOrder(tn *TreeNode) []int {
	if tn == nil {
		return nil
	}
	ans := []int{tn.val}
	ans = append(ans, preOrder(tn.left)...)
	ans = append(ans, preOrder(tn.right)...)
	return ans
}

// addList 两个链表代表的数相加
func addList(left, right *ListNode) *ListNode {
	dummp := &ListNode{}
	p := dummp

	proceed := false // 进位标志
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

// generateParenthesis 生成有效的括号组合
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	generateParenthesis0(&ans, n, n, "")
	return ans
}

// generateParenthesis0 生成括号的递归辅助函数
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

// removeNthFromEnd 删除链表倒数第n个节点
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

// makeChange 计算凑出目标金额所需的最少硬币数
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

// patternMatch 模式匹配
// pattern: 模式串
// words: 待匹配的字符串
func patternMatch(pattern, words string) bool {
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

// maxDepth 计算二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.left), maxDepth(root.right))
}

// isMirror 判断二叉树是否对称
func isMirror(node *TreeNode) bool {
	return isMirror0(node, node)
}

// isMirror0 判断对称性的递归辅助函数
func isMirror0(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.val == r.val && isMirror0(l.left, r.right) && isMirror0(l.right, r.left)
}

// minPathSum 计算网格的最小路径和
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	// 初始化第一列
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// 初始化第一行
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 动态规划填充其余位置
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

// searchRange 在排序数组中查找元素的起始和结束位置
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
			// 向左扩展
			for start-1 >= 0 && nums[start-1] == target {
				start--
			}
			// 向右扩展
			for end+1 <= len(nums)-1 && nums[end+1] == target {
				end++
			}
			return []int{start, end}
		}
	}
	return []int{-1, -1}
}

// rightSideView 获取二叉树的右视图
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

// lengthOfLongestSubstring 计算最长不含重复字符的子串长度
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

// numIslands 计算岛屿数量
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

// numIslands0 深度优先搜索标记连通的岛屿
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

// WeightSelect 根据权重随机选择
// vw: 值-权重映射表
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
