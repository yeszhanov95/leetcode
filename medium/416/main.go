package main

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum % 2 != 0 { return false }
	sum = sum / 2

	dp := make([][]bool, len(nums) + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum + 1)
		dp[i][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		num := nums[i-1]
		for j := 1; j <= sum; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
			} else if num <= j {
				dp[i][j] = dp[i-1][j-num]
			}
		}
	}

	return dp[len(nums)][sum]
}