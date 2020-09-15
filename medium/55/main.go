package main

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	minRequiredJumps := 1
	for i := len(nums)-2; i >= 0; i-- {
		if nums[i] < minRequiredJumps {
			minRequiredJumps++
		} else {
			minRequiredJumps = 1
		}
	}

	return minRequiredJumps == 1
}