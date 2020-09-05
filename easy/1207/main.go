package main

func uniqueOccurrences(arr []int) bool {
	cache1, cache2 := make(map[int]int), make(map[int]struct{})
	for _, v := range arr {
		cache1[v]++
	}
	for _, v := range cache1 {
		if _, ok := cache2[v]; ok {
			return false
		}
		cache2[v] = struct{}{}
	}
	return true
}