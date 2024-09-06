package findkpairswithsmallestsums

import (
	"container/heap"
)

// まず、 nums1[i]とnums2[0]をすべてminHeapに突っ込む
// 要素がk個埋まるまで、次を繰り返す
//   - minHeapから取り出す
//   - ペアを回答結果に突っ込む
//   - 取り出したペアの、nums2のインデックスを1増やしてminHeapに突っ込む
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	sumHeap := &sumMinHeap{}
	heap.Init(sumHeap)

	for _, n1 := range nums1 {
		// まず、候補はnums1の各要素とnums2の最初の要素の和
		p := pair{
			n1:      n1,
			n2:      nums2[0],
			n2Index: 0,
		}
		heap.Push(sumHeap, p)
	}

	ans := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		x := heap.Pop(sumHeap)
		if p, ok := x.(pair); ok {
			ans = append(ans, []int{p.n1, p.n2}) // 現在の最小ペアを追加

			// nums2のインデックスを一つ進めて、heapに登録する
			if p.n2Index+1 < len(nums2) {
				nextPair := pair{
					n1:      p.n1,
					n2:      nums2[p.n2Index+1],
					n2Index: p.n2Index + 1,
				}
				heap.Push(sumHeap, nextPair)
			}
		}
	}

	return ans
}

type pair struct {
	n1      int
	n2      int
	n2Index int
}

func (p pair) sum() int {
	return p.n1 + p.n2
}

type sumMinHeap []pair

func (s sumMinHeap) Len() int           { return len(s) }
func (s sumMinHeap) Less(i, j int) bool { return s[i].sum() < s[j].sum() }
func (s sumMinHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *sumMinHeap) Push(x any) {
	p, _ := x.(pair)
	*s = append(*s, p)
}

func (s *sumMinHeap) Pop() any {
	old := *s
	n := old.Len()
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}
