package main

import "container/heap"

func topKFrequent(words []string, k int) []string {
	cache := make(map[string]int, len(words))
	for _, word := range words {
		cache[word]++
	}

	heapArr := HeapArr{}
	heap.Init(&heapArr)

	for key, val := range cache {
		heap.Push(&heapArr, &HeapItem{key, val})
		if len(heapArr) > k {
			heap.Pop(&heapArr)
		}
	}

	result := make([]string, k)
	for i := len(heapArr) - 1; i >= 0; i-- {
		result[i] = heap.Pop(&heapArr).(*HeapItem).word
	}

	return result
}

type HeapItem struct {
	word string
	frequency int
}

type HeapArr []*HeapItem

func (h HeapArr) Len() int { return len(h) }

func (h HeapArr) Less(i, j int) bool {
	if h[i].frequency == h[j].frequency {
		return h[i].word > h[j].word
	}
	return h[i].frequency < h[j].frequency
}

func (h HeapArr) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *HeapArr) Push(x interface{}) { *h = append(*h, x.(*HeapItem)) }

func (h *HeapArr) Pop() interface{} {
	old := *h
	item := old[len(old)-1]
	*h = old[:len(old)-1]
	return item
}