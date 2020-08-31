package main

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
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

	return merge(sortList(head), sortList(rightHead))
}

func merge(left *ListNode, right *ListNode) *ListNode {
	var head *ListNode
	if left.Val > right.Val {
		head = right
		right = right.Next
	} else {
		head = left
		left = left.Next
	}

	node := head
	for left != nil && right != nil {
		if left.Val > right.Val {
			node.Next = right
			right = right.Next
		} else {
			node.Next = left
			left = left.Next
		}
		node = node.Next
	}

	if left == nil {
		node.Next = right
	} else if right == nil {
		node.Next = left
	}

	return head
}

func getMiddle(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}

	fastptr, slowptr := node.Next, node
	for fastptr != nil {
		fastptr = fastptr.Next
		if fastptr != nil {
			fastptr = fastptr.Next
			slowptr = slowptr.Next
		}
	}

	return slowptr
}