package main

type NumMatrix struct {
	matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	for i := 0; i < len(matrix); i++ {
		for j:= 1; j < len(matrix[i]); j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	return NumMatrix{matrix: matrix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int

	for ; row1 <= row2; row1++ {
		if col1 > 0 {
			sum += this.matrix[row1][col2] - this.matrix[row1][col1-1]
		} else {
			sum += this.matrix[row1][col2]
		}
	}

	return sum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */