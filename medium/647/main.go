package main

func countSubstrings(s string) int {
	runes := []rune(s)
	count := len(runes)

	for j := 0; j < len(runes)-1; j++ {
		for i := j+1; i < len(runes); i++ {
			if runes[j] == runes[i] {
				if isPalindromic(runes[j : i+1]) {
					count++
				}
			}
		}
	}

	return count
}

func isPalindromic(runes []rune) bool {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
