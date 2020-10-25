package main

func hammingDistance(x int, y int) int {
	distance, xor := 0, x ^ y
	for xor > 0 {
		distance++
		xor &= xor - 1
	}
	return distance
}