package main

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	cache := make(map[int][]int, len(nums))

	for i, v := range nums {
		cache[v] = append(cache[v], i)
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 0; i < len(nums); i++ {
		for ; i < len(nums) - 1 && nums[i] == nums[i+1]; i++ {}
		smaller := len(nums) - i - 1

		idxs := cache[nums[i]]
		for ; len(idxs) > 0; idxs = idxs[1:] {
			result[idxs[0]] = smaller
		}
	}

	return result
}