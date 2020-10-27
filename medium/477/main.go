package main

func totalHammingDistance(nums []int) int {
	var total int
	for bitPos := 0; bitPos < 32; bitPos++ {
		zeroesCount := 0
		for _, num := range nums {
			if (num & (1 << bitPos)) == 0 { zeroesCount++ }
		}
		total += zeroesCount * (len(nums) - zeroesCount)
	}
	return total
}