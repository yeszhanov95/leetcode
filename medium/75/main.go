package main

func sortColors(nums []int)  {
	idx, red, blue := 0, 0, len(nums) - 1
	for idx <= blue {
		switch nums[idx] {
		case 0:
			nums[red], nums[idx] = nums[idx], nums[red]
			red, idx = red + 1, idx + 1
		case 2:
			nums[blue], nums[idx] = nums[idx], nums[blue]
			blue--
		default: idx++
		}
	}
}