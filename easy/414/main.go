package main

func thirdMax(nums []int) int {
	min := ^(1<<32)+1
	first, second, third := min, min, min
	for _, v := range nums {
		if v > first {
			first, second, third = v, first, second
		} else if v > second && v < first {
			second, third = v, second
		} else if v > third && v < second {
			third = v
		}
	}
	if third == min {
		return first
	}
	return third
}