package main

func climbStairs(n int) int {
	cache := map[int]int{ 1 : 1, 2 : 2 }
	for i := 3; i <= n; i++ {
		if _, ok := cache[n]; ok {
			continue
		}
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n]
}