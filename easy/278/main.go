package main

func firstBadVersion(n int) int {
	l, r, m := 1, n, 0
	for l < r {
		m = l + (r - l) / 2
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func isBadVersion(n int) bool {
	return true
}