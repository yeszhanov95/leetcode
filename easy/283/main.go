package main

func moveZeroes(nums []int)  {
	idx := 0
	for _, v := range nums {
		if v != 0 {
			nums[idx] = v
			idx++
		}
	}
	for ; idx < len(nums); idx++ {
		nums[idx] = 0
	}
}