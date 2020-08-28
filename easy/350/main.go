package main

func intersect(nums1 []int, nums2 []int) []int {
	var maxLen int
	if len(nums1) > len(nums2) {
		maxLen = len(nums1)
	} else {
		maxLen = len(nums2)
	}
	cache := make(map[int]int, maxLen)
	result := make([]int, maxLen)

	for i := range nums1 {
		cache[nums1[i]]++
	}

	var idx int
	for _, v := range nums2 {
		if kV, ok := cache[v]; ok && kV > 0 {
			result[idx] = v
			idx++
			cache[v]--
		}
	}

	return result[:idx]
}