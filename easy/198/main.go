package main

func rob(nums []int) int {
	if len(nums) == 0 { return 0 }

	var dp = make([]int, len(nums)+1)
	dp[1] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i+1] = max(dp[i], nums[i] + dp[i-1])
	}

	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b { return a }
	return b
}