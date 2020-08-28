package main

func intersection(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int, len(nums1) + len(nums2))
	for i := 0; i < len(nums1); i++ {
		if _, ok := dict[nums1[i]]; !ok {
			dict[nums1[i]]++
		}
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := dict[nums2[i]]; ok {
			dict[nums2[i]]++
		}
	}

	result := make([]int, len(nums1) + len(nums2))
	var idx int
	for k, v := range dict {
		if v > 1 {
			result[idx] = k
			idx++
		}
	}

	return result[:idx]
}