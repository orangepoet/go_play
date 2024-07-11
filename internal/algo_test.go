// algorithm test
package internal

import (
	"fmt"
	"reflect"
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

func TestResolvePrime(t *testing.T) {
	testCase := map[int][]int{
		10: {2, 5}, 20: {2, 2, 5}, 30: {2, 3, 5}, 35: {5, 7},
	}
	for k, v := range testCase {
		r := resolvePrime(k)
		if !reflect.DeepEqual(r, v) {
			t.Errorf("expected %v but %v", v, r)
		}
	}
}

func TestComb(t *testing.T) {
	ret := comb(9, 2, 12)
	expected := [][]int{{3, 9}, {4, 8}, {5, 7}}
	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("expected %v but %v", expected, ret)
	}
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{s: "{{}}"}, want: true},
		{name: "test2", args: args{s: "{{}"}, want: false},
		{name: "test3", args: args{s: "}{"}, want: false},
		{name: "test3", args: args{s: "{}{"}, want: false},
		{name: "test3", args: args{s: "{}{}"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeListNode(t *testing.T) {
	type args struct {
		list []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{list: []*ListNode{makeList([]int{1, 2, 3}), makeList([]int{2, 3, 4})}}, want: makeList([]int{1, 2, 2, 3, 3, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeListNode(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preOrder(t *testing.T) {
	//     1
	//   2    5
	// 3  4  6   7
	// ->  1   2  3  4  5 6 7
	tn := &TreeNode{
		val: 1,
		left: &TreeNode{
			val: 2, left: &TreeNode{val: 3}, right: &TreeNode{val: 4},
		},
		right: &TreeNode{
			val: 5, left: &TreeNode{
				val: 6,
			}, right: &TreeNode{
				val: 7,
			},
		},
	}
	ans := preOrder(tn)
	want := []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(ans, want) {
		t.Errorf("preOrder = %v, want %v", ans, want)
	}
}

func Test_addList(t *testing.T) {
	ans := addList(makeList([]int{2, 4, 3}), makeList([]int{5, 6, 4}))
	wanted := makeList([]int{7, 0, 8})
	if !reflect.DeepEqual(ans, wanted) {
		t.Errorf("addList() = %v, want %v", ans, wanted)
	}
}

func Test_generateParenthesis(t *testing.T) {
	ans := generateParenthesis(3)
	wanted := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	if !reflect.DeepEqual(ans, wanted) {
		t.Errorf("addList() = %v, want %v", ans, wanted)
	}
}

func Test_removeNthFromEnd(t *testing.T) {
	want := makeList([]int{1, 2, 3, 4, 6})
	if got := removeNthFromEnd(makeList([]int{1, 2, 3, 4, 5, 6}), 4); !reflect.DeepEqual(got, want) {
		t.Errorf("removeNthFromEnd() = %v, want %v", got, want)
	}
}

//if got := makeChange(3000, []int{1, 7, 9}); got != 334 {
//		t.Errorf("makeChange() = %v, want %v", got, 334)
//	}
//	if got := makeChange(1, []int{1, 7, 9}); got != 1 {
//		t.Errorf("makeChange() = %v, want %v", got, 1)
//	}
//	if got := makeChange(2, []int{1, 7, 9}); got != 2 {
//		t.Errorf("makeChange() = %v, want %v", got, 2)
//	}

func Test_makeChange(t *testing.T) {
	type args struct {
		n     int
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "t1", args: args{n: 3000, coins: []int{1, 7, 9}}, want: 334},
		{name: "t2", args: args{n: 1, coins: []int{1, 7, 9}}, want: 1},
		{name: "t3", args: args{n: 2, coins: []int{1, 7, 9}}, want: 2},
		{name: "t4", args: args{n: 6, coins: []int{1, 7, 9}}, want: 6},
		{name: "t5", args: args{n: 7, coins: []int{1, 7, 9}}, want: 1},
		{name: "t6", args: args{n: 16, coins: []int{1, 7, 9}}, want: 2},
		{name: "t7", args: args{n: 17, coins: []int{1, 7, 9}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeChange(tt.args.n, tt.args.coins); got != tt.want {
				t.Errorf("makeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_patternMatch(t *testing.T) {
	match := patternMatch("abba", "dog cat cat dog")
	if !match {
		t.Errorf("patternMatch() = %v, want %v", match, true)
	}
}

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: &TreeNode{
					val:   1,
					left:  &TreeNode{val: 2},
					right: &TreeNode{val: 3},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMirror(t *testing.T) {
	r1 := &TreeNode{val: 0,
		left: &TreeNode{val: 1,
			left:  &TreeNode{val: 2},
			right: &TreeNode{val: 3},
		},
		right: &TreeNode{val: 1,
			left:  &TreeNode{val: 2},
			right: &TreeNode{val: 3},
		},
	}
	r2 := &TreeNode{val: 0,
		left: &TreeNode{val: 1,
			left:  &TreeNode{val: 2},
			right: &TreeNode{val: 3}},
		right: &TreeNode{val: 1,
			left:  &TreeNode{val: 3},
			right: &TreeNode{val: 2},
		},
	}
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{node: r1}, want: false},
		{args: args{node: r2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMirror(tt.args.node); got != tt.want {
				t.Errorf("isMirror() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{grid: [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}}, want: 7},
		{args: args{grid: [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 7
	r := searchRange(nums, target)
	fmt.Println(r)
}

func Test_RightSideView(t *testing.T) {
	//        1
	//     2     5
	//   3   4  6  7
	root := &TreeNode{
		val: 1,
		left: &TreeNode{
			val: 2,
			left: &TreeNode{
				val: 3,
			},
			right: &TreeNode{
				val: 4,
			},
		},
		right: &TreeNode{
			val: 5,
			left: &TreeNode{
				val: 6,
			},
			right: &TreeNode{
				val: 7,
			},
		},
	}
	ret := rightSideView(root)
	fmt.Println(ret)
}

func Test_LongSubString(t *testing.T) {
	l := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(l)
	if l != 3 {
		t.Errorf("length = %v, want %v", l, 3)
	}

	l2 := lengthOfLongestSubstring("bbbbbb")
	fmt.Println(l2)
	if l2 != 1 {
		t.Errorf("length = %v, want %v", l, 3)
	}
	l3 := lengthOfLongestSubstring("pwwkew")
	fmt.Println(l3)
	if l3 != 3 {
		t.Errorf("length = %v, want %v", l, 3)
	}
	l4 := lengthOfLongestSubstring(" ")
	fmt.Println(l4)
	if l4 != 1 {
		t.Errorf("length = %v, want %v", l, 3)
	}
	l5 := lengthOfLongestSubstring("au")
	if l5 != 2 {
		t.Errorf("length = %v, want %v", l, 3)

	}
}

func Test_IsLands(t *testing.T) {
	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	nums := numIslands(grid)
	if nums != 3 {
		t.Errorf("nums = %v, want %v", nums, 3)
	}
}

func Test_weightSelect(t *testing.T) {
	vw := map[int]int{1: 1, 2: 0, 3: 2, 4: 1}
	var valueMap = make(map[int]int)
	for i := 0; i < 10000; i++ {
		v, got := weightSelect(vw)
		if !got {
			fmt.Errorf("not got")
		}
		valueMap[v] = valueMap[v] + 1
	}
	fmt.Printf("%v\n", valueMap)
}
