package main

type Node struct {
	Val int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil { return [][]int{} }
	traversal := make([][]int, 0)
	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)
		subList := make([]int, 0, size)

		for ; size > 0; size-- {
			node := queue[0]
			queue = queue[1:]
			subList = append(subList, node.Val)
			if len(node.Children) > 0 {
				queue = append(queue, node.Children...)
			}
		}
		traversal = append(traversal, subList)
	}

	return traversal
}