package main

 //Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	previousNode, currentNode := head, head.Next
	for currentNode != nil {
		if currentNode.Val < previousNode.Val {
			previousNode.Next = currentNode.Next

			var previousNode2 *ListNode
			currentNode2 := head
			for currentNode2 != nil {
				if currentNode2.Val > currentNode.Val {
					currentNode.Next = currentNode2
					if previousNode2 != nil {
						previousNode2.Next = currentNode
					} else {
						head = currentNode
					}
					break
				} else {
					previousNode2 = currentNode2
					currentNode2 = currentNode2.Next
				}
			}

			currentNode = previousNode.Next
		} else {
			previousNode = currentNode
			currentNode = currentNode.Next
		}
	}

	return head
}