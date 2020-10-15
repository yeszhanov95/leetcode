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
		switch i {
		case 0: tmp = tmp[1:]
		case (len(nums) - 1): tmp = tmp[:len(nums) - 1]
		default: tmp = append(tmp[:i], tmp[i+1:]...)
		}
		helper(tmp, append(cur, nums[i]), res)
	}
}