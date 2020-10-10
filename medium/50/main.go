package main

func myPow(x float64, n int) float64 {
	switch {
	case n == 0: return 1
	case n == 1: return x
	case n < 0: return 1 / myPow(x, -n)
	}

	val := myPow(x, n/2)
	if n % 2 == 0 {
		return val * val
	}
	return val * val * x
}