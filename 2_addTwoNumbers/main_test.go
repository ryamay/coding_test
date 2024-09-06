package addtwonumbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				l1: []int{2, 4, 3},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "",
			args: args{
				l1: []int{9, 9, 9, 9, 9, 9, 9},
				l2: []int{9, 9, 9, 9},
			},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(
				intsToListNode(tt.args.l1),
				intsToListNode(tt.args.l2)); !reflect.DeepEqual(
				listNodesToInts(got), tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func intsToListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	start := &ListNode{Val: nums[0], Next: nil}

	current := start
	for i := 1; i < len(nums); i++ {
		next := &ListNode{Val: nums[i], Next: nil}
		current.Next = next
		current = next
	}

	return start
}

func listNodesToInts(node *ListNode) []int {
	ans := []int{}

	for node != nil {
		ans = append(ans, node.Val)
		node = node.Next
	}
	return ans
}
