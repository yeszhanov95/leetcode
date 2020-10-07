package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	switch {
	case root == nil: return nil
	case root.Val == val: return root
	case root.Val > val: return searchBST(root.Left, val)
	default: return searchBST(root.Right, val)
	}
}