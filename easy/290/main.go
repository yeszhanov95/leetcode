package main

func wordPattern(pattern string, str string) bool {
	cachePattern, cacheWord := make(map[byte]string, len(pattern)), make(map[string]byte, len(pattern))
	wordCnt := 0

	for i := range pattern {
		if wordCnt >= len(str) {
			return false
		}
		ptrn := pattern[i]
		startPos, hasLetter := wordCnt, false
		for ; wordCnt < len(str); wordCnt++ {
			if isSeperator(str[wordCnt]) {
				if hasLetter {
					break
				} else {
					startPos++
					continue
				}
			} else {
				hasLetter = true
			}
		}

		word := str[startPos:wordCnt]
		if v, ok := cachePattern[ptrn]; ok {
			if word != v {
				return false
			}
		} else if v2, ok2 := cacheWord[word]; ok2 {
			if ptrn != v2 {
				return false
			}
		} else {
			cachePattern[ptrn] = word
			cacheWord[word] = ptrn
		}
	}

	if wordCnt < len(str) {
		return false
	}

	return true
}

func isSeperator(b byte) bool {
	if b < 'a' || b > 'z' {
		return true
	}
	return false
}