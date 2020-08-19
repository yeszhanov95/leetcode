package main

func replaceElements(arr []int) []int {
	maxRight := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		val := arr[i]
		arr[i] = maxRight
		if val > maxRight {
			maxRight = val
		}
	}
	arr[len(arr)-1] = -1
	return arr
}