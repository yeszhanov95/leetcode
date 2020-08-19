package main

import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)
	i, j := 0, len(runes) - 1
	for i < j {
		if !unicode.IsLetter(runes[i]) && !unicode.IsNumber(runes[i]) {
			i++
			continue
		}
		if !unicode.IsLetter(runes[j]) && !unicode.IsNumber(runes[j]) {
			j--
			continue
		}
		if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
			return false
		}
		i++
		j--
	}
	return true
}