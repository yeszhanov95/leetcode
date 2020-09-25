package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := target - (nums[0] + nums[1] + nums[2])

	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		for j, k := i + 1, len(nums) - 1; j < k; {
			sum := v + nums[j] + nums[k]
			curDiff := target - sum
			minDiff = min(curDiff, minDiff)

			if sum < target {
				for j++; j < k && nums[j] == nums[j-1]; j++ {}
			} else if sum > target {
				for k--; j < k && nums[k] == nums[k+1]; k-- {}
			} else {
				return target
			}
		}
	}

	return target - minDiff
}

func min(a, b int) int {
	if abs(a) == abs(b) {
		if a < b {
			return a
		}
		return b
	}
	if abs(a) < abs(b) {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}