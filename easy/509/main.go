package main

func fib(N int) int {
	cache := make(map[int]int, N)
	return helper(N, cache)
}

func helper(n int, cache map[int]int) int {
	if n < 2 {
		return n
	} else if _, ok := cache[n]; !ok {
		cache[n] = helper(n-1, cache) + helper(n-2, cache)
	}
	return cache[n]
}