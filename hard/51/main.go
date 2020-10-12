package main

var board [][]int
var result [][]string

func solveNQueens(n int) [][]string {
	if n == 0 { return [][]string{ {} }}
	if n == 1 { return [][]string{ {"Q"} }}
	if n < 4 { return [][]string{} }

	board = make([][]int, 0, n)
	for i := 0; i < n; i++ {
		board = append(board, make([]int, n))
	}

	result = make([][]string, 0)
	helper(0, n, nil)

	return result
}

func helper(row, n int, cur []string) {
	for col := 0; col < n; col++ {
		if row == 0 { cur = make([]string, 0, n) }

		if board[row][col] == 0 {
			cur = append(cur, placeQueen(row, col))

			if row + 1 == n && len(cur) == n {
				result = append(result, cur)
			} else {
				helper(row + 1, n, append([]string{}, cur...))
			}

			cur = cur[:len(cur)-1]
			removeQueen(row, col)
		}
	}
}

func placeQueen(row, col int) string {
	for i := 0; i < len(board); i++ { board[row][i]++ }
	for i := 1; i < len(board); i++ { board[i][col]++ }

	for i, j := row + 1, col + 1; i < len(board) && j < len(board); i, j = i + 1, j + 1 { board[i][j]++ }
	for i, j := row - 1, col - 1; i >= 0 && j >= 0; i, j = i - 1, j - 1 { board[i][j]++ }
	for i, j := row - 1, col + 1; i >= 0 && j < len(board); i, j = i - 1, j + 1 { board[i][j]++ }
	for i, j := row + 1, col - 1; i < len(board) && j >= 0; i, j = i + 1, j - 1 { board[i][j]++ }

	runes := make([]rune, 0, len(board))
	for i := 0; i < len(board); i++ {
		if i == col {
			runes = append(runes, 'Q')
		} else {
			runes = append(runes, '.')
		}
	}
	return string(runes)
}

func removeQueen(row, col int) {
	for i := 0; i < len(board); i++ { board[row][i]-- }
	for i := 1; i < len(board); i++ { board[i][col]-- }

	for i, j := row + 1, col + 1; i < len(board) && j < len(board); i, j = i + 1, j + 1 { board[i][j]-- }
	for i, j := row - 1, col - 1; i >= 0 && j >= 0; i, j = i - 1, j - 1 { board[i][j]-- }
	for i, j := row - 1, col + 1; i >= 0 && j < len(board); i, j = i - 1, j + 1 { board[i][j]-- }
	for i, j := row + 1, col - 1; i < len(board) && j >= 0; i, j = i + 1, j - 1 { board[i][j]-- }
}