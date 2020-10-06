package main

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	helper(head, nil)

	return newHead
}

func helper(node, prev *ListNode) {
	if node == nil || node.Next == nil {
		return
	}

	tmp := node.Next
	node.Next = tmp.Next
	tmp.Next = node
	if prev != nil {
		prev.Next = tmp
	}

	helper(node.Next, node)
}