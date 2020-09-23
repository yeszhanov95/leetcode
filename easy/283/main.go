package main

func moveZeroes(nums []int)  {
	for i, j := 0, 1; i < len(nums) && j < len(nums); i++ {
		if nums[i] == 0 {
			for ; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					j++
					break
				}
			}
		} else {
			j++
		}
	}
}