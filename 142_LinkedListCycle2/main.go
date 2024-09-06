package linkedlistcycle2

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
func detectCycle(head *ListNode) *ListNode {
	ps := []*ListNode{}
	for head != nil {
		for _, p := range ps {
			if head == p {
				return head
			}
		}
		ps = append(ps, head)
		head = head.Next
	}
	return nil
}

// 計算すると……頭からループ先頭までと、うさぎとカメのアルゴリズムでポインタが出会う点からループ先頭までの距離は同じになるので、それを利用して解くことができる
func detectCycle2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	slow := head
	fast := head
	isCycle := false

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			isCycle = true
			break
		}
	}

	if !isCycle {
		return nil
	}

	for head != slow {
		head = head.Next
		slow = slow.Next
	}

	return head
}
