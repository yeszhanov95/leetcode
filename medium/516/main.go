package main

func longestPalindromeSubseq(s string) int {
	l := len(s)
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
		dp[i][i] = 1
	}
	for left := l - 1; left >= 0; left-- {
		for right := left + 1; right < l; right++ {
			if s[left] == s[right] {
				dp[left][right] = 2 + dp[left+1][right-1]
			} else {
				dp[left][right] = max(dp[left][right-1], dp[left+1][right])
			}
		}
	}
	return dp[0][l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}