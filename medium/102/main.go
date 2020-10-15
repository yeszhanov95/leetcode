package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	helper(root, 0)
	return res
}

func helper(node *TreeNode, depth int) {
	if node == nil { return }
	if depth == len(res) { res = append(res, []int{}) }
	res[depth] = append(res[depth], node.Val)
	helper(node.Left, depth + 1)
	helper(node.Right, depth + 1)
}