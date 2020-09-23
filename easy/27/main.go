package main

func removeElement(nums []int, val int) int {
	curIdx, lastIdx := 0, len(nums) - 1
	for curIdx <= lastIdx {
		if nums[curIdx] == val {
			nums[curIdx], lastIdx = nums[lastIdx], lastIdx - 1
		} else {
			curIdx++
		}
	}
	return curIdx
}