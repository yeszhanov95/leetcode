package main

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	if numRows == 0 {
		return result
	}
	result[0] = []int{1}
	if numRows == 1 {
		return result
	}
	result[1] = []int{ 1, 1 }

	for i := 2; i < numRows; i++ {
		inner := make([]int, i + 1)
		inner[0] = 1
		inner[len(inner)-1] = 1

		for j := 1; j < len(inner) - 1; j++ {
			inner[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = inner
	}

	return result
}