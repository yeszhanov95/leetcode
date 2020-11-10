package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil { return nil}
	helper(root)
	return root
}

func helper(node *TreeNode) {
	if node.Left == nil && node.Right == nil { return }
	if node.Left != nil { helper(node.Left) }
	if node.Right != nil { helper(node.Right) }
	node.Left, node.Right = node.Right, node.Left
}