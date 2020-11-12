package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	if root == nil { return }
	stack, node := make([]*TreeNode, 0), root
	for node != nil {
		if node.Right != nil { stack = append(stack, node.Right) }
		var next *TreeNode
		if node.Left != nil {
			next = node.Left
		} else if len(stack) > 0 {
			next, stack = stack[len(stack) - 1], stack[:len(stack) - 1]
		}
		node.Left, node.Right = nil, next
		node = next
	}
}