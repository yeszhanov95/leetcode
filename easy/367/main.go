package main

func isPerfectSquare(num int) bool {
	l, r, m := 0, num, 0
	for l <= r {
		m = (l + r) / 2
		if m * m > num {
			r = m - 1
		} else if m * m < num {
			l = m + 1
		} else {
			return true
		}
	}
	return false
}