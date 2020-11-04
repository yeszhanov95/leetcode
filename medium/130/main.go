package main

func solve(board [][]byte)  {
	if len(board) == 0 { return }
	rows, cols := len(board), len(board[0])
	for c := 0; c < cols; c++ {
		dfs(0, c, &board)
		if rows > 1 { dfs(rows - 1, c, &board) }
	}
	for r := 1; r < rows; r++ {
		dfs(r, 0, &board)
		if cols > 1 { dfs(r, cols - 1, &board) }
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Z' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(i, j int, board *[][]byte) {
	if i < 0 || j < 0 || i == len(*board) || j == len((*board)[0]) || (*board)[i][j] != 'O' {
		return
	}
	(*board)[i][j] = 'Z'
	dfs(i + 1, j, board)
	dfs(i - 1, j, board)
	dfs(i, j + 1, board)
	dfs(i, j - 1, board)
}