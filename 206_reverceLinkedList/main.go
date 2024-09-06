package revercelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 初回に思いついた実装
func reverseList(head *ListNode) *ListNode {
	reversed := (*ListNode)(nil) // ReversedListの現在位置を示すポインタ
	preReversed := &ListNode{}   // ReversedListの現在位置より一つ前に位置するポインタ（Nextフィールドだけ利用する）

	for head != nil {
		// Reversedに追加する要素を作る
		current := ListNode{
			Val:  head.Val,
			Next: reversed,
		}

		// 次のステップ用に更新
		preReversed.Next = &current // ひとつ前に
		reversed = &current         // ひとつ前に
		head = head.Next            // ひとつ後ろに
	}

	return preReversed.Next
}

// 再帰実装
func reverseList2(head *ListNode) *ListNode {
	return reverseRec(head, nil)
}

func reverseRec(head, reversed *ListNode) *ListNode {
	if head == nil {
		// headがnilの場合は、対象リスト末尾に到達したのでreversedが解
		return reversed
	}

	// headがnon-nilの場合は、reversedListの要素を追加して次を計算
	reversedHead := &ListNode{
		Val:  head.Val,
		Next: reversed,
	}
	return reverseRec(head.Next, reversedHead)
}

// 1ポインタだけで済ませるループ実装
func reverseList3(head *ListNode) *ListNode {
	reversedHead := (*ListNode)(nil)
	for head != nil {
		nextNode := head.Next    // 次のループに進むための参照を保持
		head.Next = reversedHead // Nextフィールドを逆向きに更新

		// 次のステップ用に更新
		reversedHead = head
		head = nextNode
	}

	return reversedHead
}
