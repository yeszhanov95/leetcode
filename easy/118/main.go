package main

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for idx := 0 ; numRows > 0; numRows, idx = numRows-1, idx+1 {
		slice := make([]int, idx+1)
		slice[0], slice[len(slice)-1] = 1, 1
		for j := 1; j < len(slice)-1; j++ {
			slice[j] = res[idx-1][j-1] + res[idx-1][j]
		}
		res[idx] = slice
	}
	return res
}