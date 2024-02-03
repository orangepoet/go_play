// algorithm test
package algo

import (
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
