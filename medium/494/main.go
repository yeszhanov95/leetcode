package main

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	cache := make(map[int]map[int]int, len(nums))
	return dfs(nums, 0, 0, S, cache)
}

func dfs(nums []int, idx int, curr, target int, cache map[int]map[int]int) int {
	if idx == len(nums) {
		if curr == target { return 1 }
		return 0
	}

	if _, ok := cache[idx]; ok {
		if _, ok2 := cache[idx][curr]; ok2 {
			return cache[idx][curr]
		}
	} else {
		cache[idx] = make(map[int]int)
	}

	res := dfs(nums, idx + 1, curr + nums[idx], target, cache) + dfs(nums, idx + 1, curr - nums[idx], target, cache)
	cache[idx][curr] = res

	return res
}