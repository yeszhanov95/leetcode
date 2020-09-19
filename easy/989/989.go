package main

import "strconv"

func addToArrayForm(A []int, K int) []int {
	res := make([]int, len(A)+1)
	S := strconv.Itoa(K)
	if len(S) > len(A) {
		res = make([]int, len(S)+1)
	}
	res[0] = -1

	for carry, i, j, k := 0, len(A)-1, len(S)-1, len(res)-1; k >= 0; carry, i, j, k = carry/10, i-1, j-1, k-1 {
		if i >= 0 {
			carry += A[i]
		}
		if j >= 0 {
			carry += int(S[j] - '0')
		}
		res[k] = carry % 10
	}

	if res[0] == 0 {
		return res[1:]
	}
	return res
}