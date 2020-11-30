package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for row := range dp {
		dp[row] = make([]int, n)
		dp[row][0] = 1
	}
	for col := 1; col < n; col++ {
		dp[0][col] = 1
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[row][col] = dp[row][col-1] + dp[row-1][col]
		}
	}

	return dp[m-1][n-1]
}