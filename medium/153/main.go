package main

func findMin(nums []int) int {
	l, r, m := 0, len(nums) - 1, 0
	for l < r {
		m = (r - l) / 2 + l
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}