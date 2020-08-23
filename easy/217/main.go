package main

func containsDuplicate(nums []int) bool {
	tmp := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := tmp[v]; ok {
			return true
		}
		tmp[v] = struct{}{}
	}
	return false
}