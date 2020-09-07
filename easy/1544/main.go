package main

func makeGood(s string) string {
	stack := make([]rune, 0, len(s))
	for _, r := range s {
		if len(stack) != 0 && stack[len(stack)-1] != r && toLower(stack[len(stack)-1]) == toLower(r) {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, r)
	}
	return string(stack)
}

//using two pointers sln
func makeGood1(s string) string {
	j := 0
	b := []rune(s)

	for i := range b {
		b[j] = b[i]
		if j > 0 && b[j] != b[j-1] && toLower(b[j]) == toLower(b[j-1]) {
			j -= 2
		}
		j++
	}

	return string(b[:j])
}

func toLower(r rune) rune {
	if r < 'a' {
		return r + 32
	}
	return r
}

