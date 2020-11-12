package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil { return 0 }
	var totalSum int
	dfs(root, 0, &totalSum)
	return totalSum
}

func dfs(node *TreeNode, curSum int, totalSum *int) {
	curSum = (curSum * 10) + node.Val
	if node.Left == nil && node.Right == nil { *totalSum += curSum }
	if node.Left != nil { dfs(node.Left, curSum, totalSum) }
	if node.Right != nil { dfs(node.Right, curSum, totalSum) }
}