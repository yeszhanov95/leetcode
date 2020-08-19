package main

func sortedSquares(A []int) []int {
	length := len(A)
	res := make([]int, length)
	i, j := 0, length -1

	for z := length - 1; z >= 0; z-- {
		sqI, sqJ := A[i] * A[i], A[j] * A[j]
		if sqI > sqJ {
			res[z] = sqI
			i++
		} else {
			res[z] = sqJ
			j--
		}
	}

	return res
}