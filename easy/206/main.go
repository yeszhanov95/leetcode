package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	rightHead := slow.Next
	slow.Next = nil
	return reverse(reverseList(head), reverseList(rightHead))
}

func reverse(left *ListNode, right *ListNode) *ListNode {
	node := right
	for node.Next != nil {
		node = node.Next
	}
	node.Next = left
	return right
}