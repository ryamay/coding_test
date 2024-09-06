package linkedlistcycle2

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	lists := []*ListNode{
		{
			Val: 3,
		},
		{
			Val: 2,
		},
		{
			Val: 0,
		},
		{
			Val: -4,
		},
	}
	lists[0].Next = lists[1]
	lists[1].Next = lists[2]
	lists[2].Next = lists[3]
	lists[3].Next = lists[1]

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{lists[0]},
			want: lists[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
