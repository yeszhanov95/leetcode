package main


 type ListNode struct {
 	Val int
 	Next *ListNode
 }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var head *ListNode
	if l1.Val > l2.Val {
		head, l2 = l2, l2.Next
	} else {
		head, l1 = l1, l1.Next
	}

	node := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			node.Next, l2 = l2, l2.Next
		} else {
			node.Next, l1 = l1, l1.Next
		}
		node = node.Next
	}

	if l1 == nil {
		node.Next, l2 = l2, l2.Next
	} else {
		node.Next, l1 = l1, l1.Next
	}

	return head
}