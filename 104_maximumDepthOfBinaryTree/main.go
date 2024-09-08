package maximumdepthofbinarytree

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
func maxDepth(root *TreeNode) int {
	return maxDepthRec(root)
}

func maxDepthRec(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(maxDepthRec(node.Left), maxDepthRec(node.Right))
}
