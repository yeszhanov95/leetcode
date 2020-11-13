package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil { return [][]int{} }
	result, stack := make([][]int, 0, 1), make([]*TreeNode, 0, 1)
	stack = append(stack, root)
	var toRight = true
	for len(stack) > 0 {
		size := len(stack)
		subList := make([]int, 0, size)
		for i := 0; i < size; i++ {
			var node  = stack[size-i-1]
			subList = append(subList, node.Val)
		}

		for j := size - 1; j >= 0; j-- {
			node := stack[j]
			if toRight {
				if node.Left != nil { stack = append(stack, node.Left) }
				if node.Right != nil { stack = append(stack, node.Right) }
			} else {
				if node.Right != nil { stack = append(stack, node.Right) }
				if node.Left != nil { stack = append(stack, node.Left) }
			}

		}

		stack = stack[size:]
		result = append(result, subList)
		toRight = !toRight
	}
	return result
}