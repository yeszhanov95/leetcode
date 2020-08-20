package main

func findMaxAverage(nums []int, k int) float64 {
	max, current := 0, 0

	for i := 0; i < k; i++ {
		current += nums[i]
	}
	max = current

	for i:= k; i < len(nums); i++ {
		current = current + nums[i] - nums[i-k]
		if current > max {
			max = current
		}
	}

	return float64(max) / float64(k)
}