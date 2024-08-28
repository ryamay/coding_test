package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	vals := []int{}
	for head != nil {
		currentVal := head.Val
		for v := range vals {
			if currentVal == v {
				return true // サイクル検出
			}
		}
		vals = append(vals, currentVal)
		head = head.Next
	}
	return false
}

// ウサギとカメのアルゴリズムを使う解法（ベスト）
func hasCycle2(head *ListNode) bool {
	fastP := head // fastは2つずつ進む
	for fastP != nil && fastP.Next != nil {
		head = head.Next // headは1つずつ進む
		if head == nil {
			return false // 終端したら、サイクルなし
		}

		fastP = fastP.Next.Next
		if head == fastP {
			return true // fastPがheadに追いついたら、それはサイクルがあるということ
		}
	}
	return false
}
