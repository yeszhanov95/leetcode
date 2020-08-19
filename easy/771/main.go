package main

func numJewelsInStones(J string, S string) int {
	var stones int
	for _, vS := range S {
		for _, vJ := range J {
			if vS == vJ {
				stones++
				break
			}
		}
	}
	return stones
}