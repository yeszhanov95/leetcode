package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	prefix := make([]byte, 0)
	minWordLength := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minWordLength {
			minWordLength = len(strs[i])
		}
	}

	for i := 0; i < minWordLength; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[j-1][i] {
				return string(prefix)
			}
		}
		prefix = append(prefix, strs[0][i])
	}
	return string(prefix)
}