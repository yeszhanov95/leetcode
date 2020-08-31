package main

import "fmt"

func main() {
	slice := []int{ 2, 5, -1, 8, 6, 9 }
	fmt.Println(slice)
	sorted := heapSort(slice)
	fmt.Println(sorted)
}

func heapSort(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)

	buildMaxHeap(sorted)

	for i := len(sorted) - 1; i >= 1; i-- {
		sorted[0], sorted[i] = sorted[i], sorted[0]
		heapifyDown(sorted, 0, i)
	}

	return sorted
}

func buildMaxHeap(slice []int) {
	for i := (len(slice) - 1) / 2; i >= 0; i-- {
		heapifyDown(slice, i, len(slice))
	}
}

func heapifyDown(slice []int, root, hi int) {
	for {
		leftIdx := (root * 2) + 1
		if leftIdx >= hi {
			break
		}

		maxIdx := leftIdx
		if rightIdx := leftIdx + 1; rightIdx < hi && slice[rightIdx] > slice[leftIdx] {
			maxIdx = rightIdx
		}

		if slice[root] >= slice[maxIdx] {
			break
		}

		slice[root], slice[maxIdx] = slice[maxIdx], slice[root]
		root = maxIdx
	}
}