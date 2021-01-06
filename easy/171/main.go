package main

func titleToNumber(s string) int {
	columnNumber, n := 0, 1
	for i := len(s) - 1; i >= 0; i-- {
		columnNumber += int(s[i] - 'A' + 1) * n
		n *= 26
	}
	return columnNumber
}