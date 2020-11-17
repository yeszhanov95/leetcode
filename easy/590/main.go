package main

type Node struct {
	Val int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil { return []int{} }
	traversal := make([]int, 0)
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		traversal = append(traversal, node.Val)
		if len(node.Children) > 0 {
			stack = append(stack, node.Children...)
		}
	}

	for i, j := 0, len(traversal)-1; i < j; i, j = i+1, j-1 {
		traversal[i], traversal[j] = traversal[j], traversal[i]
	}

	return traversal
}