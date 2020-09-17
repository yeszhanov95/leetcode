package main

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)

	length := len(words)
	if length == 0 {
		return ""
	} else if length == 1 {
		return words[0]
	}

	slice := make([]rune, 0)
	for i := length - 1; i >= 0; i-- {
		slice = append(slice, []rune(words[i])...)
		slice = append(slice, ' ')
	}

	return string(slice[:len(slice)-1])
}