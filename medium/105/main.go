package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inMap := make(map[int]int, len(inorder))
	for i := range inorder {
		inMap[inorder[i]] = i
	}
	var preStart = 0
	return helper(preorder, &preStart, inMap, 0, len(inorder) - 1)
}

func helper(preorder []int, preStart *int, inMap map[int]int, inStart, inEnd int) *TreeNode {
	if inStart > inEnd || *preStart == len(preorder) { return nil }

	node := &TreeNode{ Val:preorder[*preStart] }
	*preStart++
	inRoot := inMap[node.Val]

	node.Left = helper(preorder, preStart, inMap, inStart, inRoot - 1)
	node.Right = helper(preorder, preStart, inMap, inRoot + 1, inEnd)

	return node
}