package main

var queenPlacements []int
var total, N int

func totalNQueens(n int) int {
	if n == 1 { return 1 }
	if n < 4 { return 0 }

	queenPlacements, total, N = make([]int, 0 , n), 0 , n
	helper(0)

	return total
}

func helper(row int) {
	for col := 0; col < N; col++ {
		if isValid(row, col) {
			queenPlacements = append(queenPlacements, col)
			if row + 1 == N {
				total++
			} else {
				helper(row + 1)
			}
			queenPlacements = queenPlacements[:len(queenPlacements)-1]
		}
	}
}

func isValid(row, col int) bool {
	for i := 0; i < row; i++ {
		diff := abs(queenPlacements[i] - col)
		if diff == 0 || diff == row - i {
			return false
		}
	}
	return true
}

func abs(num int) int {
	if num < 0 { return -num }
	return num
}