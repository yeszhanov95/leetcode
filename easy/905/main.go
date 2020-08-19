package main

func sortArrayByParity(A []int) []int {
	idx := 0
	for i, v := range A {
		if v % 2 == 0 {
			A[idx], A[i] = v, A[idx]
			idx++
		}
	}
	return A
}