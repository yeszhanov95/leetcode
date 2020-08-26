package main

import "fmt"

func main() {
	arr := []int{ 2, 5, -1, 8, 6, 9 }
	fmt.Println(arr)
	sorted := insertionSort(arr)
	fmt.Println(sorted)
}

func insertionSort(arr []int) []int {
	length := len(arr)
	result := make([]int, length)
	copy(result, arr)

	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if result[j] >= result[j-1] {
				break
			}
			result[j], result[j-1] = result[j-1], result[j]
		}
	}

	return result
}