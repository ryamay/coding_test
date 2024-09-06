package removeduplicatesfromsortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
昇順にソート済みのため、次の要素との値を比較して、
値が一致した場合に head.Nextをhead.Next.Nextに置き換える
ことにすれば重複を削除できる

n要素に対して、
計算量はO(n),
メモリ消費量はO(1)
*/
func deleteDuplicates(head *ListNode) *ListNode {
	start := head // listの先頭要素を返すので、保持しておく

	for head != nil {
		if head.Next != nil {
			// 昇順にソート済みのため、次の要素との値を比較して、
			// 値が一致した場合に head.Nextをhead.Next.Nextに置き換える
			// ことにすれば重複を削除できる
			if head.Val == head.Next.Val {
				head.Next = head.Next.Next
				continue
			}
		}

		head = head.Next
	}

	return start
}
