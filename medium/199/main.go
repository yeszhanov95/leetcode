package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil { return []int{} }
	rightSide := make([]int, 0)
	helper(root, &rightSide, 0)
	return rightSide
}

func helper(node *TreeNode, rightSide *[]int, depth int) {
	if len(*rightSide) == depth { *rightSide = append(*rightSide, node.Val) }
	if node.Right != nil { helper(node.Right, rightSide, depth + 1) }
	if node.Left != nil { helper(node.Left, rightSide, depth + 1) }
}