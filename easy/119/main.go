package main

func getRow(rowIndex int) []int {
	res, cache := make([]int, rowIndex+1), make(map[int]map[int]int)
	res[0], res[rowIndex] = 1, 1

	for i := 1; i < rowIndex; i++ {
		res[i] = getCellValue(rowIndex, i, cache)
	}

	return res
}

func getCellValue(row, col int, cache map[int]map[int]int) int {
	if col == 0 || row == 1 || col == row {
		return 1
	}
	if _, ok := cache[row]; !ok {
		cache[row] = make(map[int]int)
	}
	if _, ok := cache[row][row-col]; ok {
		return cache[row][row-col]
	}
	if _, ok := cache[row][col]; !ok {
		cache[row][col] = getCellValue(row-1, col-1, cache) + getCellValue(row-1, col, cache)
	}
	return cache[row][col]
}