package convertsortedarraytobinarysearchtree

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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	nodeIdx := len(nums) / 2

	node := &TreeNode{Val: nums[nodeIdx]}
	node.Left = sortedArrayToBST(nums[:nodeIdx])
	node.Right = sortedArrayToBST(nums[nodeIdx+1:])

	return node
}
