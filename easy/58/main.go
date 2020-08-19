package main

func lengthOfLastWord(s string) int {
	var length int
	for i := len(s) - 1; i > -1; i-- {
		if s[i] != ' ' {
			length++
		} else if (i<len(s)-1) && s[i+1] != ' '{
			break
		}
	}
	return length
}