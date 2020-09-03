package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	result := make([]int, k)
	cache := make(map[int]int, len(nums))

	for _, v := range nums {
		cache[v]++
	}

	h := &IntHeap{}
	heap.Init(h)

	for key, val := range cache {
		if len(*h) == k {
			if (*h)[0].freq > val {
				continue
			}
			heap.Pop(h)
		}
		heap.Push(h, MyStruct{num: key, freq: val})
	}

	idx := 0
	for _, v := range *h {
		result[idx] = v.num
		idx++
	}

	return result
}

type MyStruct struct {
	num int
	freq int
}

type IntHeap []MyStruct

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(MyStruct))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}