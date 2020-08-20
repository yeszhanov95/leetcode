package main

func missingNumber(nums []int) int {
	expectedSum, actualSum := len(nums), 0
	for i, v := range nums {
		actualSum += v
		expectedSum += i
	}
	return expectedSum - actualSum
}