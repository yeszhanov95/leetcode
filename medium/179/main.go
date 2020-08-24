package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	length := len(nums)
	if length == 0 {
		return ""
	} else if length == 1 {
		return strconv.Itoa(nums[0])
	}

	sort.Slice(nums, func(i, j int) bool {
		first, second := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return first+second >= second+first
	})

	if nums[0] == 0 {
		return "0"
	}

	var str strings.Builder
	for _, v := range nums {
		str.WriteString(strconv.Itoa(v))
	}
	return str.String()
}