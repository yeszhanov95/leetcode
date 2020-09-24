package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	} else if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	res, btm, top := make([]string, 0), nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] - nums[i-1] == 1 {
			top = nums[i]
		} else {
			if btm == top {
				res = append(res, strconv.Itoa(btm))
			} else {
				res = append(res, fmt.Sprintf("%v->%v", btm, top))
			}
			btm, top = nums[i], nums[i]
		}
	}

	if btm == top {
		res = append(res, strconv.Itoa(btm))
	} else {
		res = append(res, fmt.Sprintf("%v->%v", btm, top))
	}

	return res
}