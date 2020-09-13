package main

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, cur := head, head.Next
	for ; cur != nil; cur = cur.Next {
		if prev.Val == cur.Val {
			prev.Next = nil
		} else {
			prev.Next = cur
			prev = cur
		}
	}

	return head
}
