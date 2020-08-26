package main

import "fmt"

func main() {
	arr := []int{ 2, 5, 2, 8, 6, 9 }
	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j >= 0; j-- {
			if arr[j] >= arr[j-1] {
				break
			}
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}