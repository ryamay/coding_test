package mergesortedarray

import (
	"reflect"
	"testing"
)

func Test_merge_mとnの和が1(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "m<n",
			args: args{nums1: []int{0}, m: 0, nums2: []int{2}, n: 1},
			want: []int{2},
		},
		{
			name: "m>n",
			args: args{nums1: []int{3}, m: 1, nums2: []int{}, n: 0},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("got %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}

func Test_merge_mとnの和が2(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "m<n",
			args: args{nums1: []int{0, 0}, m: 0, nums2: []int{2, 3}, n: 2},
			want: []int{2, 3},
		},
		{
			name: "m=n=1, nums1[0] < nums2[0]",
			args: args{nums1: []int{1, 0}, m: 1, nums2: []int{2}, n: 1},
			want: []int{1, 2},
		},
		{
			name: "m=n=1, nums1[0] > nums2[0]",
			args: args{nums1: []int{3, 0}, m: 1, nums2: []int{2}, n: 1},
			want: []int{2, 3},
		},
		{
			name: "m>n",
			args: args{nums1: []int{4, 5}, m: 2, nums2: []int{}, n: 0},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("got %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}

func Test_merge_mとnの和が3(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "m=0, n=3",
			args: args{nums1: []int{0, 0, 0}, m: 0, nums2: []int{2, 3, 3}, n: 3},
			want: []int{2, 3, 3},
		},
		{
			name: "m=1, n=2, 末尾に挿入",
			args: args{nums1: []int{3, 0, 0}, m: 1, nums2: []int{2, 2}, n: 2},
			want: []int{2, 2, 3},
		},
		{
			name: "m=1, n=2, 真ん中に挿入",
			args: args{nums1: []int{3, 0, 0}, m: 1, nums2: []int{2, 4}, n: 2},
			want: []int{2, 3, 4},
		},
		{
			name: "m=1, n=2, 先頭に挿入",
			args: args{nums1: []int{3, 0, 0}, m: 1, nums2: []int{4, 5}, n: 2},
			want: []int{3, 4, 5},
		},
		{
			name: "m=2, n=1, 末尾に挿入",
			args: args{nums1: []int{2, 2, 0}, m: 2, nums2: []int{3}, n: 1},
			want: []int{2, 2, 3},
		},
		{
			name: "m=2, n=1, 真ん中に挿入",
			args: args{nums1: []int{2, 5, 0}, m: 2, nums2: []int{1}, n: 1},
			want: []int{1, 2, 5},
		},
		{
			name: "m=2, n=1, 先頭に挿入",
			args: args{nums1: []int{2, 2, 0}, m: 2, nums2: []int{1}, n: 1},
			want: []int{1, 2, 2},
		},
		{
			name: "m=3, n=0",
			args: args{nums1: []int{4, 5, 5}, m: 3, nums2: []int{}, n: 0},
			want: []int{4, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("got %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
