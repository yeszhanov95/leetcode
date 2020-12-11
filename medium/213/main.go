package main

func rob(nums []int) int {
	if len(nums) == 1 { return nums[0] }
	return max(helper(nums[:len(nums)-1]), helper(nums[1:]))
}

func helper(houses []int) int {
	prev2, prev1 := 0, houses[0] // prev2 = i - 2, prev = i - 1
	for i := 1; i < len(houses); i++ {
		tmp := prev1
		prev1 = max(tmp, prev2 + houses[i])
		prev2 = tmp
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}