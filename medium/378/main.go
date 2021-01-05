package main

import "container/heap"

func kthSmallest(matrix [][]int, k int) int {
	h := make(IntHeap, 0, len(matrix) * len(matrix))
	for row := range matrix {
		for col := range matrix[row] {
			h = append(h, matrix[row][col])
		}
	}
	heap.Init(&h)
	var res interface{}
	for k > 0 {
		res = heap.Pop(&h)
		k--
	}
	return res.(int)
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}