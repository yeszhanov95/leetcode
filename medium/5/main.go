package main

func longestPalindrome(s string) string {
	var longestPalindrome = s[:1]
	var palindromes = make(map[string]int)
	for i := 1; i < len(s); i++ {
		var char = s[i]
		for j := 0; j < i; j++ {
			if char == s[j] && ((i - j) < 3 || palindromes[s[j+1:i]] > 0) {
				palindromes[s[j:i+1]] = i - j + 1
				break
			}
		}
	}
	for palindrome, length := range palindromes {
		if length > len(longestPalindrome) { longestPalindrome = palindrome }
	}
	return longestPalindrome
}