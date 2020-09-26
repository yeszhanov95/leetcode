package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int, 0)

	for left, leftVal := range nums {
		if left > 0 && leftVal == nums[left-1] {
			continue
		}

		for right := length - 1; (left + 2) < right; right-- {
			if right != length - 1 && nums[right] == nums[right + 1] {
				continue
			}
			innerLeft, innerRight := left + 1, right - 1

			for innerLeft < innerRight {
				sum := leftVal + nums[innerLeft] + nums[innerRight] + nums[right]
				if sum < target {
					for innerLeft++; innerLeft < innerRight && nums[innerLeft] == nums[innerLeft-1]; innerLeft++ {}

				} else if sum > target {
					for innerRight--; innerLeft < innerRight && nums[innerRight] == nums[innerRight+1]; innerRight-- {}
				} else {
					res = append(res, []int{leftVal, nums[innerLeft], nums[innerRight], nums[right]})
					for innerLeft++; innerLeft < innerRight && nums[innerLeft] == nums[innerLeft-1]; innerLeft++ {}
					for innerRight--; innerLeft < innerRight && nums[innerRight] == nums[innerRight+1]; innerRight-- {}
				}
			}
		}
	}

	return res
}