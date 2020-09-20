package main

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	node := res
	for carry := 0; l1 != nil || l2 != nil || carry > 0; carry = carry/10 {
		if l1 != nil {
			carry, l1 = carry + l1.Val, l1.Next
		}
		if l2 != nil {
			carry, l2 = carry + l2.Val, l2.Next
		}
		node.Next = &ListNode{Val: carry % 10}
		node = node.Next
	}
	return res.Next
}