package main

import "strings"

func uncommonFromSentences(A string, B string) []string {
	words := strings.Split(A, " ")
	words = append(words, strings.Split(B, " ")...)

	cache :=  make(map[string]int, len(words))
	for _, word := range words {
		cache[word]++
	}

	res := make([]string, 0)
	for word, cnt := range cache {
		if cnt == 1 {
			res = append(res, word)
		}
	}

	return res
}