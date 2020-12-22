package main

func exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}
	for row := range board {
		for col := range board[row] {
			if helper(board, used, word, 0, row, col) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, used [][]bool, word string, i, row, col int) bool {
	if i < len(word) && board[row][col] == word[i] && !used[row][col] {
		i++
		used[row][col] = true
		m, n := len(board), len(board[0])
		if col + 1 < n && helper(board, used, word, i, row, col + 1) {
			return true
		}
		if col - 1 >= 0 && helper(board, used, word, i, row, col - 1) {
			return true
		}
		if row + 1 < m && helper(board, used, word, i, row + 1, col) {
			return true
		}
		if row - 1 >= 0 && helper(board, used, word, i, row - 1, col) {
			return true
		}
		used[row][col] = false
	}

	if i >= len(word) {
		return true
	}
	return false
}
