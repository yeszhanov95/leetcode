package main

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int, len(nums))
	for i := range nums {
		if i2, ok := cache[target - nums[i]]; ok && i != i2 {
			return []int{i, i2}
		}
		cache[nums[i]] = i
	}

	return []int{}
}