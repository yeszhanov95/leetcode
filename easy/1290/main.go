package main

import "math"

type ListNode struct {
	Val int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	slice := make([]float64, 0, 30)
	for head != nil {
		slice = append(slice, float64(head.Val))
		head = head.Next
	}

	var res float64
	for i, pow := len(slice) - 1, 0.0; i >=0; i, pow = i-1, pow+1 {
		if slice[i] == 1 {
			res += math.Pow(2.0, pow)
		}
	}

	return int(res)
}
