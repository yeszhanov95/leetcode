package main

import "unicode"

func letterCasePermutation(S string) []string {
	lettersPos := make([]int, 0)
	for i, r := range S {
		if unicode.IsLetter(r) { lettersPos = append(lettersPos, i) }
	}

	result := make([]string, 0, 1 << len(lettersPos))
	for bitMask := 0; bitMask < cap(result); bitMask++ {
		runes := []rune(S)

		for k := 0; k < len(lettersPos); k++ {
			if (bitMask & (1 << k)) != 0 {
				letterPos := lettersPos[k]
				if unicode.IsUpper(runes[letterPos]) {
					runes[letterPos] = unicode.ToLower(runes[letterPos])
				} else {
					runes[letterPos] = unicode.ToUpper(runes[letterPos])
				}
			}
		}

		result = append(result, string(runes))
	}

	return result
}