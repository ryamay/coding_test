package removeduplicatesfromsortedlist2

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "",
			args: []int{1, 2, 3, 3, 4, 4, 5},
			want: []int{1, 2, 5},
		},
		{
			name: "",
			args: []int{1, 1, 1, 2, 3},
			want: []int{2, 3},
		},
		{
			name: "",
			args: []int{1, 1, 1, 2, 2},
			want: []int{},
		},
		{
			name: "",
			args: []int{1, 2, 3, 3, 3},
			want: []int{1, 2},
		},
		{
			name: "",
			args: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "",
			args: []int{1, 1},
			want: []int{},
		},
		{
			name: "",
			args: []int{1},
			want: []int{1},
		},
		{
			name: "",
			args: []int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(intsToListNode(tt.args)); !reflect.DeepEqual(listNodesToInts(got), tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", listNodesToInts(got), tt.want)
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
