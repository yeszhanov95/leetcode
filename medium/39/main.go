package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	helper(candidates, 0, []int{}, target, &result)
	return result
}

func helper(nums []int, i int, current []int, remainder int, result *[][]int) {
	if remainder == 0 {
		*result = append(*result, append([]int{}, current...))
		return
	}
	if remainder < 0 || i == len(nums) || nums[i] > remainder {
		return
	}

	helper(nums, i, append(current, nums[i]), remainder - nums[i], result)
	helper(nums, i + 1, current, remainder, result)
}