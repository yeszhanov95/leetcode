package main

func longestCommonSubsequence(text1 string, text2 string) int {
	length1, length2 := len(text1), len(text2)
	if length1 < length2 {
		return longestCommonSubsequence(text2, text1)
	}

	dp1, dp2 := make([]int, length1 + 1), make([]int, length1 + 1)
	for row := 1; row <= length1; row++ {
		for col := 1; col <= length2; col++ {
			if text1[row-1] == text2[col-1] {
				dp2[col] = 1 + dp1[col-1]
			} else {
				dp2[col] = max(dp2[col-1], dp1[col])
			}
		}
		copy(dp1, dp2)
	}

	return dp2[length2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}