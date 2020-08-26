package main

import "fmt"

func main() {
	arr := []int{ 2, 5, -1, 8, 6, 9 }
	fmt.Println(arr)
	sorted := selectionSort(arr)
	fmt.Println(sorted)
}

func selectionSort(arr []int) []int {
	length := len(arr)
	result := make([]int, length)
	copy(result, arr)
	for i := 0; i < length; i++ {
		minIdx := i
		for j := i+1; j < length; j++ {
			if result[j] < result[minIdx] {
				minIdx = j
			}
		}
		if i != minIdx {
			result[i], result[minIdx] = result[minIdx], result[i]
		}
	}
	return result
}