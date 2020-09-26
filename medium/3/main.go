package main

func lengthOfLongestSubstring(s string) int {
	var max, current, start int
	cache := make(map[byte]int)

	for i := range s {
		if prevIdx, ok := cache[s[i]]; ok && prevIdx >= start {
			max = maxF(max, current)
			current, start = i - (prevIdx + 1), prevIdx + 1
		}

		cache[s[i]], current = i, current + 1
	}

	return maxF(max, current)
}

func maxF(a, b int) int {
	if a > b {
		return a
	}
	return b
}