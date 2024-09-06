package revercelinkedlist

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			args: []int{1, 2},
			want: []int{2, 1},
		},
		{
			args: []int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := reverseList(intsToListNode(tt.args)); !reflect.DeepEqual(listNodesToInts(got), tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_reverseListRec(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			args: []int{1, 2},
			want: []int{2, 1},
		},
		{
			args: []int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := reverseList2(intsToListNode(tt.args)); !reflect.DeepEqual(listNodesToInts(got), tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_reverseList3(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			args: []int{1, 2},
			want: []int{2, 1},
		},
		{
			args: []int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := reverseList3(intsToListNode(tt.args)); !reflect.DeepEqual(listNodesToInts(got), tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
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
