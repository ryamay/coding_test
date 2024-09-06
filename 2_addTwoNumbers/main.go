package addtwonumbers

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil // non-emptyなのでこのチェックは不要だが、念の為
	}

	dummy := ListNode{}
	preStart := &dummy
	head := &dummy

	var increase int // 繰り上がりを保持する変数
	for l1 != nil || l2 != nil || increase != 0 {

		sum := valOrZeroIfNil(l1) + valOrZeroIfNil(l2) + increase
		val := sum % 10 // 1の位を取得する
		increase = sum / 10

		// 計算結果を代入して、ひとつ先に進める
		head.Next = &ListNode{Val: val, Next: nil}

		head = head.Next
		l1 = nextNodeIfExists(l1)
		l2 = nextNodeIfExists(l2)
	}

	return preStart.Next
}

func valOrZeroIfNil(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}

func nextNodeIfExists(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	return node.Next
}
