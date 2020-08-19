package main

func peakIndexInMountainArray(A []int) int {
	for i := 1; i < len(A); i++ {
		if (A[i-1] < A[i]) && (A[i+1] < A[i]) {
			return i
		}
	}
	return -1
}