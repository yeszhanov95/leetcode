package main

func reverseVowels(s string) string {
	vowels := map[rune]struct{} {
		'e' : {}, 'u' : {}, 'i' : {}, 'o' : {}, 'a' : {},
		'E' : {}, 'U' : {}, 'I' : {}, 'O' : {}, 'A' : {},
	}
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for ; i < j; i++ {
			if _, ok := vowels[runes[i]]; ok {
				break
			}
		}
		for ; i < j; j-- {
			if _, ok := vowels[runes[j]]; ok {
				break
			}
		}
		if i < j {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}
	return string(runes)
}