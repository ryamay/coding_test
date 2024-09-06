package kthlargestelement

import "container/heap"

type KthLargest struct {
	k       int
	minHeap *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	// heapを作成
	minHeap := &IntHeap{}
	heap.Init(minHeap)
	kthL := KthLargest{
		k:       k,
		minHeap: minHeap,
	}
	for n := range nums {
		kthL.Add(n)
	}
	return kthL

}

func (kthL *KthLargest) Add(val int) int {
	if kthL.minHeap.Len() < kthL.k {
		heap.Push(kthL.minHeap, val)
	} else if (*kthL.minHeap)[0] < val {
		heap.Pop(kthL.minHeap)
		heap.Push(kthL.minHeap, val)
	}
	return (*kthL.minHeap)[0]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(n any) {
	*h = append(*h, n.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := h.Len()
	popped := old[n-1]
	*h = old[:n-1]
	return popped
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
