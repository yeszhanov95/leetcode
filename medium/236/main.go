package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil { return nil }
	if root == p || root == q { return root }

	leftResult := lowestCommonAncestor(root.Left, p, q)
	rightResult := lowestCommonAncestor(root.Right, p, q)

	if leftResult == nil { return rightResult }
	if rightResult == nil { return leftResult }

	return root
}