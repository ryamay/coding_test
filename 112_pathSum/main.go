package pathsum

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
func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumNode(root, targetSum)
}

func hasPathSumNode(node *TreeNode, targetSum int) bool {
	if node == nil {
		return false // nodeがnilの場合、falseを返す ※leaf判定・求める経路の判定は親nodeで行う
	}

	// leafにたどり着いたら、targetSumとnodeの値が等しいとき解となる
	if node.Left == nil && node.Right == nil {
		return targetSum == node.Val
	}

	// 左右の子の結果を確認し、どちらかに解があればtrue
	leftResult := hasPathSumNode(node.Left, targetSum-node.Val)
	rightResult := hasPathSumNode(node.Right, targetSum-node.Val)
	return leftResult || rightResult
}
