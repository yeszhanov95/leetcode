package main

func containsNearbyDuplicate(nums []int, k int) bool {
	cache := make(map[int]int, len(nums))

	for i, n := range nums {
		if _, ok := cache[n]; ok && i - cache[n] <= k {
			return true
		}
		cache[n] = i
	}

	return false
}