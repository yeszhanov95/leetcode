package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil { return [][]int{} }
	result := make([][]int, 0)
	dfs(root, []int{}, &result, sum)
	return result
}

func dfs(node *TreeNode, pathValues []int, result *[][]int, remaining int) {
	pathValues, remaining = append(pathValues, node.Val), remaining - node.Val
	if remaining == 0 && node.Left == nil && node.Right == nil {
		*result = append(*result, append(make([]int, 0, len(pathValues)), pathValues...))
		return
	}
	if node.Left != nil {
		dfs(node.Left, pathValues, result, remaining)
	}
	if node.Right != nil {
		dfs(node.Right, pathValues, result, remaining)
	}
}