package main

func search(nums []int, target int) int {
	left, right, mid := 0, len(nums) - 1, 0
	for left <= right {
		mid = left + (right - left) / 2
		if nums[mid] == target { return mid }
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}