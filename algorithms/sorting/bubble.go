package main

import "fmt"

func main() {
	arr := []int{2, 9, -1, 5, 6, 0}
	fmt.Println(arr)
	sorted := bubbleSort(arr)
	fmt.Println(sorted)
}

func bubbleSort(arr []int) []int {
	length := len(arr)
	result := make([]int, length)
	copy(result, arr)

	for i := 1; i < length; i++ {
		for j := 1; j < length; j++ {
			if result[j] < result[j-1] {
				result[j], result[j-1] = result[j-1], result[j]
			}
		}
	}

	return result
}