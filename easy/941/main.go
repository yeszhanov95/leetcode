package main

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	growth := false
	fall := false
	for i := 1; i < len(A); i++ {
		if (A[i] == A[i-1]) || (fall && A[i] > A[i-1]) || (!growth && A[i] < A[i-1]) {
			return false
		}
		if A[i] > A[i-1] {
			growth = true
		} else {
			fall = true
		}
	}
	return growth && fall
}