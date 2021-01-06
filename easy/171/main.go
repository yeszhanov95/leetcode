package main

func titleToNumber(s string) int {
	n := len(s)
	var columnNumber int
	for i := 0; i < n - 1; i++ {
		left := int(s[i] - 'A') + 1
		right := 1
		for j := n - i - 1; j > 0; j-- {
			right *= 26
		}
		columnNumber += left * right
	}
	columnNumber += int(s[n-1] - 'A') + 1
	return columnNumber
}