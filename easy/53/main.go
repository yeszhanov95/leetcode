package main

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	}

	max, current := nums[0], nums[0]
	for i := 1; i < length; i++ {
		current = getMax(nums[i], nums[i] + current)
		max = getMax(max, current)
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}