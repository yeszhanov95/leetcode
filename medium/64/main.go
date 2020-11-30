package main

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	
	dp[0] = grid[0][0]
	for col := 1; col < n; col++ {
		dp[col] = grid[0][col] + dp[col-1]
	}

	for row := 1; row < m; row++ {
		dp[0] += grid[row][0]
		for col := 1; col < n; col++ {
			dp[col] = grid[row][col] + min(dp[col], dp[col-1])
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}