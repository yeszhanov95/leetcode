package main

import "container/heap"

func findKthLargest(nums []int, k int) int {
	heapArr := make(HeapArr, len(nums))
	copy(heapArr, nums)
	heap.Init(&heapArr)
	var res int
	for ; k > 0; k-- {
		res = heap.Pop(&heapArr).(int)
	}
	return res
}

type HeapArr []int

func (h HeapArr) Len() int { return len(h) }

func (h HeapArr) Less(i, j int) bool { return h[i] > h[j] }

func (h HeapArr) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *HeapArr) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *HeapArr) Pop() interface{} {
	old := *h
	last := old[len(old)-1]
	*h = old[:len(old)-1]
	return last
}