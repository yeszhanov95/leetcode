package main

import "fmt"

func main() {
	slice := []int{8, 6, -2, -4, 0, 3}
	fmt.Println(slice)
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	quickSort(sorted, 0, len(sorted) - 1)
	fmt.Println(sorted)
}

func quickSort(slice []int, lowIdx, highIdx int) {
	if lowIdx < highIdx {
		pi := partition(slice, lowIdx, highIdx)
		quickSort(slice, lowIdx, pi - 1)
		quickSort(slice, pi + 1, highIdx)
	}
}

func partition(slice []int, lowIdx, highIdx int) int {
	pivot, idx := slice[highIdx], lowIdx
	for i := lowIdx; i < highIdx; i++ {
		if slice[i] < pivot {
			slice[idx], slice[i] = slice[i], slice[idx]
			idx++
		}
	}
	if idx != highIdx {
		slice[idx], slice[highIdx] = slice[highIdx], slice[idx]
	}
	return idx
}