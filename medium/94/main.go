package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil { return []int{} }
	result, node := make([]int, 0), root
	stack := make([]*TreeNode, 0)
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		result = append(result, node.Val)
		node = node.Right
	}
	return result
}