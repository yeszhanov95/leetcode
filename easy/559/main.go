package main

type Node struct {
	Val int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil { return 0 }
	var max int
	for _, child := range root.Children {
		cur := maxDepth(child)
		if cur > max { max = cur }
	}
	return max + 1
}