package main

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil { return []int{} }
	traversal := make([]int, 0)
	stack, node := make([]*Node, 0), root
	for {
		traversal = append(traversal, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
		if len(stack) > 0 {
			node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		} else {
			break
		}
	}
	return traversal
}