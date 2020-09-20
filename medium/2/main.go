package main

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	slice := make([]int, 0)
	for carry := 0; l1 != nil || l2 != nil || carry > 0; carry = carry/10 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		slice = append(slice, carry % 10)
	}

	res := &ListNode{}
	node := res
	for _, v := range slice {
		node.Next = &ListNode{Val:v}
		node = node.Next
	}

	return res.Next
}