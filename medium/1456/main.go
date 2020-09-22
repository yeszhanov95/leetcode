package main

func maxVowels(s string, k int) int {
	runes := []rune(s)
	vowels := map[rune]struct{} {
		'a' : struct{}{},
		'e' : struct{}{},
		'i' : struct{}{},
		'o' : struct{}{},
		'u' : struct{}{},
	}
	var  max, cur, idx int

	for  ; k > 0; idx, k = idx+1, k-1 {
		if _, ok := vowels[runes[idx]]; ok {
			cur++
		}
	}
	max = cur

	for j := 0 ; idx < len(runes); idx, j = idx+1, j+1 {
		if _, ok := vowels[runes[idx]]; ok {
			cur++
		}
		if _, ok := vowels[runes[j]]; ok {
			cur--
		}
		if cur > max {
			max = cur
		}
	}

	return max
}