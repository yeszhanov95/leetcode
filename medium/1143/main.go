package main

func longestCommonSubsequence(text1 string, text2 string) int {
	length1, length2 := len(text1), len(text2)
	if length1 == 0 || length2 == 0 {
		return 0
	}

	dp := make([][]int, length1 + 1)
	dp[0] = make([]int, length2 + 1)

	for row := 1; row <= length1; row++ {
		dp[row] = make([]int, length2 + 1)
		dp[row][0] = 0
		for col := 1; col <= length2; col++ {
			if text1[row-1] == text2[col-1] {
				dp[row][col] = 1 + dp[row-1][col-1]
			} else {
				dp[row][col] = max(dp[row-1][col], dp[row][col-1])
			}
		}
	}

	return dp[length1][length2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}