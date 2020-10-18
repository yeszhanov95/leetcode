package main

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums) - 1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	if left > right { return - 1 }
	mid := left + (right - left) / 2
	if nums[mid] < target {
		return binarySearch(nums, mid + 1, right, target)
	} else if nums[mid] > target {
		return binarySearch(nums, left, mid - 1, target)
	}
	return mid
}