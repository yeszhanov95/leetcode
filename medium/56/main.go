package main

import "sort"

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})

	result := make([][]int, length)
	idx := 0

	for i := 0; i < length-1; i++ {
		result[idx] = intervals[i]

		for j := i; j < length-1; j++ {
			if result[idx][1] >= intervals[j+1][0] && result[idx][1] <= intervals[j+1][1] {
				if intervals[j+1][1] > result[idx][1] {
					result[idx][1] = intervals[j+1][1]
				}
				i++
			} else if result[idx][1] >= intervals[j+1][1] {
				i++
				continue
			} else {
				break
			}
		}

		idx++

		if i == length-2 {
			result[idx] = intervals[i+1]
			idx++
		}


	}

	return result[:idx]
}