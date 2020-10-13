package main

var result [][]string
var queenPlacemets []int
var N int

func solveNQueens(n int) [][]string {
	if n == 0 { return [][]string{ {} }}
	if n == 1 { return [][]string{ {"Q"} }}
	if n < 4 { return [][]string{} }

	result, queenPlacemets, N = make([][]string, 0), make([]int, 0, n), n
	helper(0)
	return result
}

func helper(row int) {
	for col := 0; col < N; col++ {
		if isValid(row, col) {
			queenPlacemets = append(queenPlacemets, col)
			if row + 1 == N {
				generateCombination()
			} else {
				helper(row + 1)
			}
			queenPlacemets = queenPlacemets[:len(queenPlacemets)-1]
		}
	}
}

func isValid(row, col int) bool {
	for i := 0; i < row; i++ {
		diff := abs(queenPlacemets[i] - col)
		if diff == 0 || diff == row - i {
			return false
		}
	}
	return true
}

func generateCombination() {
	tmp := make([]string, 0, N)
	for _, col := range queenPlacemets {
		runes := make([]rune, 0, N)
		for i := 0; i < N; i++ {
			if i == col {
				runes = append(runes, 'Q')
			} else {
				runes = append(runes, '.')
			}
		}
		tmp = append(tmp, string(runes))
	}
	result = append(result, tmp)
}

func abs(num int) int {
	if num < 0 { return -num }
	return num
}