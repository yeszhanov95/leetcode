package main

import "sort"

func heightChecker(heights []int) int {
	res := 0
	ordered := make([]int, len(heights))
	copy(ordered, heights)
	sort.Ints(ordered)
	for i, v := range heights {
		if v != ordered[i] {
			res++
		}
	}
	return res
}