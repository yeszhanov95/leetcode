package main

func firstUniqChar(s string) int {
	cache := make(map[byte]int, len(s))
	for i := range s {
		cache[s[i]]++
	}
	for i := range s {
		if cache[s[i]] == 1 { return i }
	}
	return -1
}