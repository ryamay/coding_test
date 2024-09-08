package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 幅優先探索で、フェーズ別に[]intを作成して回答をつくる
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}

	q := queue{}
	q.push(root)
	for len(q) > 0 {
		// その時点でqueueに格納されているnodeを処理
		levelVals := q.processLevel()
		ans = append(ans, levelVals)
	}

	return ans
}

// 末尾にpush, 先頭からpop
type queue []*TreeNode

func (q *queue) processLevel() []int {
	levelVals := []int{}
	for range *q {
		node := q.pop()
		if node != nil {
			levelVals = append(levelVals, node.Val)

			// 次のlevel用のnodeをpush(次のレベルのnodeがすべてnilなら、queueが空になり停止)
			if node.Left != nil {
				q.push(node.Left)
			}
			if node.Right != nil {
				q.push(node.Right)
			}
		}
	}

	return levelVals
}

func (q *queue) push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *queue) pop() *TreeNode {
	node := (*q)[0]
	*q = (*q)[1:]
	return node
}
