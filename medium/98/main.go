package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(node, down, top *TreeNode) bool {
	if node == nil { return true }

	if down != nil && node.Val <= down.Val { return false }
	if top != nil && node.Val >= top.Val { return false }

	return helper(node.Left, down, node) && helper(node.Right, node, top)
}