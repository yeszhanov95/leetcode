package main

import "math"

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	stack := make([]*ListNode, 0)
	for lastVal := math.MinInt32; head != nil; head = head.Next {
		if head.Val != lastVal {
			lastVal = head.Val
			stack = append(stack, &ListNode{Val:head.Val})
			if len(stack) > 1 {
				stack[len(stack)-2].Next = stack[len(stack)-1]
			}
		} else if len(stack) > 0 && stack[len(stack)-1].Val == lastVal {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				stack[len(stack)-1].Next = nil
			}
		}
	}

	if len(stack) == 0 {
		return nil
	}
	return stack[0]
}
