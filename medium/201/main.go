package main

func rangeBitwiseAnd(m int, n int) int {
	var i int
	for m != n {
		m, n, i = m >> 1, n >> 1, i + 1
	}
	return m << i
}