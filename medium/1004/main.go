package main

func longestOnes(A []int, K int) int {
	max, cur := 0, 0

	for i, j := 0, 0; i < len(A); i++ {
		if A[i] == 1 {
			cur++
		} else if K > 0 {
			cur, K = cur + 1, K - 1
		} else {
			K--
		}

		if K < 0 {
			if A[j] == 0 {
				K++
			} else {
				cur--
			}
			j++
		}

		if cur > max {
			max = cur
		}
	}

	return max
}