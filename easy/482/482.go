package main

func licenseKeyFormatting(S string, K int) string {
	runes, k := make([]rune, 0), K

	for i := len(S)-1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		}
		if k == 0 {
			runes = append(runes, '-')
			k = K
		}
		runes = append(runes, ToUpper(rune(S[i])))
		k--
	}

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes[:len(runes)])
}

func ToUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}