package main

func updateMatrix(matrix [][]int) [][]int {
	rows, cols := make([]int, 0), make([]int, 0)
	visited := make([][]bool, 0, len(matrix))
	for i := range matrix {
		visited = append(visited, make([]bool, len(matrix[0])))
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	vertical := []int{-1, 1, 0, 0}
	horizontal := []int{0, 0, -1, 1}
	for len(rows) > 0 { // or len(cols). it doesn't matter
		i, j := rows[0], cols[0]
		rows, cols = rows[1:], cols[1:]
		isZero := matrix[i][j] == 0

		for k := 0; k < 4; k++ {
			ik, jk := i + vertical[k], j + horizontal[k]
			if ik < 0 || jk < 0 || ik == len(matrix) || jk == len(matrix[0]) || visited[ik][jk] {
				continue
			}

			visited[ik][jk] = true
			rows = append(rows, ik)
			cols = append(cols, jk)

			if !isZero && matrix[ik][jk] > 0 {
				matrix[ik][jk] = matrix[i][j] + 1
			}
		}
	}

	return matrix
}