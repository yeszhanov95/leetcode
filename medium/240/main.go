package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 { return false }
	return helper(matrix, 0, len(matrix[0]) - 1, 0, len(matrix) - 1, target)
}

func helper(matrix [][]int, left, right, top, down, target int) bool {
	if top > down || left > right { return false }
	col, row := left + (right - left) / 2, top
	for row <= down {
		if matrix[row][col] == target { return true }
		if matrix[row][col] > target { break }
		row++
	}
	return helper(matrix, left, col - 1, row, down, target) || helper(matrix, col + 1, right, top, row - 1, target)
}