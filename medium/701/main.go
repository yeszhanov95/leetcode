package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil { return &TreeNode{Val: val}}
	helper(root, val)
	return root
}

func helper(root *TreeNode, val int) {
	if val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
			return
		}
		helper(root.Right, val)
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
			return
		}
		helper(root.Left, val)
	}
}