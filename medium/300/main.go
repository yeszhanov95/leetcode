package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	var lis = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		num := nums[i]
		for j := 0; j < i; j++ {
			if num > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		lis = max(lis, dp[i])
	}
	return lis
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}