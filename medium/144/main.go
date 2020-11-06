package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil { return []int{} }
	result, node := make([]int, 0), root
	stack := make([]*TreeNode, 0)
	for {
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			node = node.Left
		} else if len(stack) > 0 {
			node, stack = stack[len(stack) - 1], stack[:len(stack) - 1]
		} else {
			break
		}
	}
	return result
}