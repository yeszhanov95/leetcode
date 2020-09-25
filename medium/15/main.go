package main

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	res, sum := make([][]int, 0), 0
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		j, k := i + 1, len(nums) - 1

		for j < k {
			sum = nums[j] + nums[k] + v
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				res = append(res, []int{v, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {}
				for k--; j < k && nums[k] == nums[k+1]; k-- {}
			}
		}
	}

	return res
}