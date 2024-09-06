package topkfrequentelement

import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	frequencyMap := map[int]int{} // key: num, value: frequency
	for _, n := range nums {
		frequencyMap[n] += 1
	}

	fh := &FrequencyHeap{}
	heap.Init(fh)
	for value, count := range frequencyMap {
		heap.Push(fh, Frequency{value: value, count: count})
		if fh.Len() > k {
			heap.Pop(fh) // 要素数がkを超えた場合、最小要素を削除する
		}
	}

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		fr := heap.Pop(fh)
		if f, ok := fr.(Frequency); ok {
			ans[i] = f.value
		}
	}
	return ans
}

type Frequency struct {
	value int
	count int
}

type FrequencyHeap []Frequency

func (h FrequencyHeap) Len() int           { return len(h) }
func (h FrequencyHeap) Less(i, j int) bool { return h[i].count < h[j].count } // 頻度順にソートし、最も頻度の低いものがpopされる
func (h FrequencyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FrequencyHeap) Push(x any) {
	*h = append(*h, x.(Frequency))
}

func (h *FrequencyHeap) Pop() (x any) {
	old := *h
	n := old.Len()
	x = old[n-1]
	*h = old[0 : n-1]
	return x
}
