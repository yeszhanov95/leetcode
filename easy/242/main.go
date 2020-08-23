package main

func isAnagram(s string, t string) bool {
	len1:=len(s)
	len2:=len(t)
	if len1 != len2 {
		return false
	}
	var s1,s2 [26]int
	for i := 0; i < len1; i++ {
		s1[s[i]-'a']++
		s2[t[i]-'a']++
	}
	return s1 == s2
}