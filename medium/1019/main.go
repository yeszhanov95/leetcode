package main

type ListNode struct {
	Val int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	nodes := make([]int, 0)
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}

	stack, result := make([][2]int, 0), make([]int, len(nodes))
	for i := 0; i < len(nodes); {
		if len(stack) == 0 || nodes[i] <= stack[len(stack)-1][1] {
			stack = append(stack, [2]int{i, nodes[i]})
			i++
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[last[0]] = nodes[i]
		}
	}

	return result
}
