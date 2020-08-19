package main

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	nums2 := make([]int, n)
	for i := 0; i < n; i++ {
		nums2[i] = i + 1
	}
	for i := range nums {
		t := nums[i]
		if nums2[t-1] != 0 {
			nums2[t-1] = 0
			n--
		}
	}
	res := make([]int, n)
	idx := 0
	for i := range nums2 {
		if nums2[i] != 0 {
			res[idx] = nums2[i]
			idx++
		}
	}
	return res
}