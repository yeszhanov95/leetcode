package main

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }
	if n == 1 { return nums[0] }

	ans := nums[0]
	prevMax, prevMin := ans, ans
	curMax, curMin := ans, ans

	for i := 1; i < n; i++ {
		num := nums[i]
		curMax = max(num, max(prevMax * num, prevMin * num))
		curMin = min(num, min(prevMax * num, prevMin * num))
		ans = max(ans, curMax)
		prevMax, prevMin = curMax, curMin
	}

	return ans
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func min(a, b int) int {
	if a < b { return a }
	return b
}