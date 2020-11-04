package main

func solve(board [][]byte)  {
	visited := make([][]bool, 0, len(board))
	for i := len(board); i > 0; i-- {
		visited = append(visited, make([]bool, len(board[0])))
	}

	var rows, cols []int
	var onBorder bool
	for i := range board {
		for j := range board[i] {
			if visited[i][j] || board[i][j] == 'X' { continue }
			rows, cols = make([]int, 0), make([]int, 0)
			onBorder = false
			dfs(i, j, board, visited, &rows, &cols, &onBorder)

			if !onBorder && len(rows) > 0 {
				for k := 0; k < len(rows); k++ {
					board[rows[k]][cols[k]] = 'X'
				}
			}
		}
	}
}

func dfs(i, j int, board [][]byte, visited [][]bool, rows, cols *[]int, onBorder *bool) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || board[i][j] == 'X' || visited[i][j] {
		return
	}

	visited[i][j] = true
	*rows, *cols = append(*rows, i), append(*cols, j)
	if !(*onBorder) && (i == 0 || j == 0 || i == len(board) - 1 || j == len(board[0]) - 1) {
		*onBorder = true
	}

	dfs(i + 1, j, board, visited, rows, cols, onBorder)
	dfs(i - 1, j, board, visited, rows, cols, onBorder)
	dfs(i, j + 1, board, visited, rows, cols, onBorder)
	dfs(i, j - 1, board, visited, rows, cols, onBorder)
}