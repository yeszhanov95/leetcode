package main

import "fmt"

func main() {
	slice := []int{2, 9, -1, 5, 6, 0}
	fmt.Println(slice)
	sorted := mergeSort(slice)
	fmt.Println(sorted)
}

func mergeSort(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	subSort(result, 0, len(slice) - 1)
	return result
}

func subSort(slice []int, leftStart, rightEnd int) {
	fmt.Println("subSort:", leftStart, rightEnd)
	if leftStart >= rightEnd {
		return
	}

	middle := (leftStart + rightEnd) / 2
	subSort(slice, leftStart, middle)
	subSort(slice, middle + 1, rightEnd)

	merge(slice, leftStart, rightEnd)
}

func merge(slice []int, leftStart, rightEnd int) {
	fmt.Println("merge:", leftStart, rightEnd)
	tmp := make([]int, len(slice))
	copy(tmp, slice)

	leftEnd := (leftStart + rightEnd) / 2
	rightStart := leftEnd + 1

	left := leftStart
	right := rightStart
	index := leftStart

	for left <= leftEnd && right <= rightEnd {
		if slice[left] > slice[right] {
			tmp[index] = slice[right]
			right++
		} else {
			tmp[index] = slice[left]
			left++
		}
		index++
	}

	copy(tmp[index : rightEnd+1], slice[right : rightEnd + 1])
	copy(tmp[index : rightEnd+1], slice[left : leftEnd + 1])

	copy(slice, tmp)
}