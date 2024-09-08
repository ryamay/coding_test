package minimumdepthofbinarytree

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
func minDepth(root *TreeNode) int {
	return minDepthRec(root)
}

func minDepthRec(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left, right := minDepthRec(node.Left), minDepthRec(node.Right)
	if left == 0 || right == 0 {
		// 少なくとも片方の子がない場合、もう片方のmindepthを足す
		// min(left, right)を使うとゼロが回答になってしまうので、
		// 存在する子のdepthを回答とするため和を返す
		return 1 + left + right
	}
	return 1 + min(left, right)
}
