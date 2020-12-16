package main

func jump(nums []int) int {
	curr, last := 0, len(nums) - 1
	if last == 0 {
		return 0
	}

	jumps := 1
	for {
		left, right := curr, curr + nums[curr]
		if right >= last {
			break
		}

		for j := left + 1; j <= right; j++ {
			if j + nums[j] > curr + nums[curr] {
				curr = j
			}
		}

		jumps++
	}

	return jumps
}