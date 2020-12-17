package main

func partitionLabels(S string) []int {
	rightmost := make([]int, 26)
	for i := range S {
		rightmost[S[i]-'a'] = i
	}

	result := make([]int, 0)
	left, right := 0, 0
	for i := range S {
		right = max(right, rightmost[S[i]-'a'])
		if i == right {
			result = append(result, right - left + 1)
			left = i + 1
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}