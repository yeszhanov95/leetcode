package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 { return false }
	m, n := len(matrix), len(matrix[0])
	return binarySearch(matrix, 1, m * n, target, n)
}

func binarySearch(matrix [][]int, left, right, target, n int) bool {
	if left > right { return false }
	mid := left + (right - left) / 2

	var i, j int
	if mid % n == 0 {
		i, j = (mid / n) - 1, n - 1
	} else {
		i, j = mid / n, (mid % n) - 1
	}

	if matrix[i][j] == target {
		return true
	}
	if matrix[i][j] > target {
		return binarySearch(matrix, left, mid - 1, target, n)
	}
	return binarySearch(matrix, mid + 1, right, target, n)
}