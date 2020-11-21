package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil { return nil }
	if root.Val == key { return deleteTargetNode(root) }

	var curr = root
	var prev *TreeNode
	for curr != nil && curr.Val != key {
		prev = curr
		if curr.Val > key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	if curr != nil {
		if prev.Left == curr {
			prev.Left = deleteTargetNode(curr)
		} else {
			prev.Right = deleteTargetNode(curr)
		}
	}

	return root
}

func deleteTargetNode(node *TreeNode) *TreeNode {
	if node.Left == nil { return node.Right }
	if node.Right == nil { return node.Left}

	next := node.Right
	var prev *TreeNode
	for next.Left != nil {
		prev = next
		next = next.Left
	}

	next.Left = node.Left
	if node.Right != next {
		prev.Left = next.Right
		next.Right = node.Right
	}

	return next
}