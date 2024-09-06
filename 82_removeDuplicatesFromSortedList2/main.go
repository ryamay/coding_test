package removeduplicatesfromsortedlist2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 昇順ソートされたLinked Listの中から、数字が重複する要素をすべて削除する
// ex.) [1,2,2,3,3,4,4,5] -> [1,2,5]
// ex.) [1,1,1,2,3] -> [2,3]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := ListNode{}
	preStart := &dummy
	distinct := &dummy

	for head != nil && head.Next != nil {
		if head.Val != head.Next.Val {
			// 重複なしなので、そのまま進める
			distinct.Next = head
			distinct = head
			head = head.Next
		} else {
			// 重複している要素をスキップする
			currentVal := head.Val
			for head.Next != nil && head.Next.Val == currentVal {
				head = head.Next // 別の値がでてくるまでスキップ
			}

			// このタイミングで、 head.Nextはnil or 別の値なので、つなぎ変える
			distinct.Next = head.Next
			head = head.Next
		}
	}

	return preStart.Next
}
