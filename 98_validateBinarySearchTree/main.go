package validatebinarysearchtree

import "math"

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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	isValidNode := maxOfTree(root.Left) < root.Val && root.Val < minOfTree(root.Right) // サブツリーのmax, minを計算しているのは効率が悪い

	return isValidNode && isValidBST(root.Left) && isValidBST(root.Right)
}

func maxOfTree(node *TreeNode) int {
	if node == nil {
		return math.MinInt
	}

	return max(node.Val, maxOfTree(node.Left), maxOfTree(node.Right))
}

func minOfTree(node *TreeNode) int {
	if node == nil {
		return math.MaxInt
	}

	return min(node.Val, minOfTree(node.Left), minOfTree(node.Right))
}

// max, minを再帰関数で保持して、計算量を減らしたバージョン
func isValidBST2(root *TreeNode) bool {
	return isBST(root, math.MinInt, math.MaxInt)
}

func isBST(node *TreeNode, treeMin, treeMax int) bool {
	if node == nil {
		return true
	}

	if node.Val <= treeMin || treeMax <= node.Val {
		return false
	}

	return isBST(node.Left, treeMin, node.Val) && isBST(node.Right, node.Val, treeMax)
}
