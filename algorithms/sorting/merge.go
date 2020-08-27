package main

import "fmt"

func main() {
	slice := []int{2, 9, -1, 0, -3, 32, 4, -843}
	fmt.Println(slice)
	sorted := mergeSort(slice)
	fmt.Println(sorted)
}

func mergeSort(slice []int) []int {
	length := len(slice)
	if length < 2 {
		return slice
	}

	m := length / 2
	return merge(mergeSort(slice[:m]), mergeSort(slice[m:]))
}

func merge(leftSlice, rightSlice []int) []int {
	leftLen, rightLen := len(leftSlice), len(rightSlice)
	size, leftIdx, rightIdx := leftLen + rightLen, 0, 0
	slice := make([]int, size)

	for k := 0; k < size; k++ {
		if leftIdx > leftLen - 1 {
			slice[k] = rightSlice[rightIdx]
			rightIdx++
		} else if rightIdx > rightLen - 1 {
			slice[k] = leftSlice[leftIdx]
			leftIdx++
		} else if leftSlice[leftIdx] < rightSlice[rightIdx] {
			slice[k] = leftSlice[leftIdx]
			leftIdx++
		} else {
			slice[k] = rightSlice[rightIdx]
			rightIdx++
		}
	}

	return slice
}