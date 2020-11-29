package main

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins) + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount + 1)
		dp[i][0] = 1
	}
	for i := 1; i <= len(coins); i++ {
		coin := coins[i-1]
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= coin {
				dp[i][j] += dp[i][j-coin]
			}
		}
	}
	return dp[len(coins)][amount]
}