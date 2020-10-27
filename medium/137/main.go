package main

func singleNumber(nums []int) int {
	var result int
	for bitPos := 0; bitPos < 64; bitPos++ {
		bitMask, onesCount := 1 << bitPos, 0
		for _, num := range nums {
			if (num & bitMask) != 0 { onesCount++ }
		}
		if (onesCount % 3) == 1 {
			result |= bitMask
		}
	}
	return result
}