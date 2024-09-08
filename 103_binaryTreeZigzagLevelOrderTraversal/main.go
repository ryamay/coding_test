package binarytreezigzaglevelordertraversal

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}

	s := stack{}
	s.push(root)

	level := 0
	for len(s) > 0 {
		level++
		levelVals := s.processLevel(level)
		ans = append(ans, levelVals)
	}

	return ans
}

// 末尾に追加し、末尾から取り出す
// stack構造を採用したが、push時の対処が複雑になった。
// queue構造を使ったうえで、各levelごとにlevelValsのappend方式を変えるだけの方が単純そう。
type stack []*TreeNode

func (s *stack) processLevel(level int) []int {

	// 現在レベルのnodeを取り出す
	levelNodes := []*TreeNode{}
	for range *s {
		nd := s.pop()
		levelNodes = append(levelNodes, nd)
	}

	levelVals := []int{}
	for _, n := range levelNodes {
		levelVals = append(levelVals, n.Val)
		s.pushNextLevel(n, level) // 次レベルのnodeをpushしておく
	}

	return levelVals
}

func (s *stack) pushNextLevel(node *TreeNode, level int) {
	leftToRight := level%2 == 1 // 奇数/偶数によって、push順序を変える

	if leftToRight {
		if node.Left != nil {
			s.push(node.Left)
		}
		if node.Right != nil {
			s.push(node.Right)
		}
	} else {
		if node.Right != nil {
			s.push(node.Right)
		}
		if node.Left != nil {
			s.push(node.Left)
		}
	}
}

func (s *stack) push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *stack) pop() *TreeNode {
	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return node
}
