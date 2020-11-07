package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil { return []int{} }
	stack, result := make([]*TreeNode, 0), make([]int, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}