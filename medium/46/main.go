package main

func permute(nums []int) [][]int {
	res, cur := make([][]int, 0), make([]int, 0, len(nums))
	helper(nums, cur, &res)
	return res
}

func helper(nums, cur []int, res *[][]int) {
	if len(cur) == cap(cur) {
		*res = append(*res, append([]int{}, cur...))
		return
	}
	for i := 0; i < len(nums); i++ {
		tmp := append([]int{}, nums...)
		helper(append(tmp[:i], tmp[i+1:]...), append(cur, nums[i]), res)
	}
}