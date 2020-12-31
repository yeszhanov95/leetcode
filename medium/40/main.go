package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
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
	if remainder < 0 || i == len(nums) || remainder < nums[i] {
		return
	}

	helper(nums, i + 1, append(current, nums[i]), remainder - nums[i], result)
	for i++; i < len(nums) && nums[i] == nums[i-1]; i++ { }
	if i < len(nums) {
		helper(nums, i, current, remainder, result)
	}
}