package main

func twoSum(numbers []int, target int) []int {
	cache := make(map[int]int)
	for i, v := range numbers {
		if _, ok := cache[v]; ok {
			return []int{cache[v], i+1}
		}
		cache[target-v] = i+1
	}
	return []int{}
}