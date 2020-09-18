package main

func reverseWords(s string) string {
	runes := []rune(s)
	for i, j := 0, 0; i < len(runes); i++ {
		if s[i] == ' ' || i == len(runes) - 1 {
			k := i - 1
			if i == len(runes) - 1 {
				k = i
			}
			for ; j < k; j, k = j + 1, k - 1 {
				runes[j], runes[k] = runes[k], runes[j]
			}
			j = i + 1
		}
	}
	return string(runes)
}