package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := helper(nums, target, 0, len(nums), true)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	return []int{ left, helper(nums, target, left + 1, len(nums), false) - 1 }
}

func helper(nums []int, target, l, r int, isLeft bool) int {
	for l < r {
		m := (r - l) / 2 + l
		if nums[m] > target || (isLeft && nums[m] == target) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}