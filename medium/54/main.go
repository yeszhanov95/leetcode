package main

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return []int{}
	} else if rows == 1 {
		return matrix[0]
	}
	cols := len(matrix[0])
	totalItems := rows * cols
	result := make([]int, totalItems)

	i, j, idx, diff := 0, 0, 0, 0

	for idx < totalItems {
		for ; j < cols; j++ {
			result[idx] = matrix[i][j]
			idx++
		}
		j--
		cols--
		if idx == len(result) {
			return result
		}

		for i++ ; i < rows; i++ {
			result[idx] = matrix[i][j]
			idx++
		}
		i--
		rows--
		if idx == len(result) {
			return result
		}

		for j--; j >= diff; j-- {
			result[idx] = matrix[i][j]
			idx++
		}
		j++
		if idx == len(result) {
			return result
		}

		for i--; i > diff; i-- {
			result[idx] = matrix[i][j]
			idx++
		}
		i++

		j++
		diff++
	}

	return result
}