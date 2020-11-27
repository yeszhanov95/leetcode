package main

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(dp); i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}
	return min(dp[len(dp)-1], dp[len(dp)-2])
}

func min(a, b int) int  {
	if a < b { return a }
	return b
}