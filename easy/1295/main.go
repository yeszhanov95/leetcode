package main

import "fmt"

func findNumbers(nums []int) (res int) {
	for i := range nums {
		if len(fmt.Sprintf("%d", nums[i])) % 2 == 0 {
			res++
		}
	}
	return res
}