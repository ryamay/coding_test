package mergetwobinarytrees

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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// root1, root2を同時にトレースしていく。
	// 両方がnilならnilを設定、
	// 上記以外の場合はroot1, root2の和を設定する（nilは0扱い）
	return mergeTreesRec(root1, root2)
}

func mergeTreesRec(node1, node2 *TreeNode) *TreeNode {
	if node1 == nil && node2 == nil {
		return nil
	}
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	newNode := &TreeNode{Val: node1.Val + node2.Val}
	newNode.Left = mergeTreesRec(node1.Left, node2.Left)
	newNode.Right = mergeTreesRec(node1.Right, node2.Right)
	return newNode
}
