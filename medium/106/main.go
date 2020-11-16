package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	inMap := make(map[int]int, len(inorder))
	for i := range inorder {
		inMap[inorder[i]] = i
	}
	var postEnd = len(postorder) - 1
	return helper(postorder, &postEnd, inMap, 0, len(inorder) - 1)
}

func helper(postorder []int, postEnd *int, inMap map[int]int, inStart, inEnd int) *TreeNode {
	if inStart > inEnd || *postEnd < 0 { return nil }

	node := &TreeNode{Val:postorder[*postEnd]}
	*postEnd--;
	inRoot := inMap[node.Val]

	node.Right = helper(postorder, postEnd, inMap, inRoot + 1, inEnd)
	node.Left = helper(postorder, postEnd, inMap, inStart, inRoot - 1)

	return node
}