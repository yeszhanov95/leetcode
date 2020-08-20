package main

func findWords(words []string) []string {
	keyboard := make(map[byte]int, 26)

	keyboard['q'] = 1
	keyboard['w'] = 1
	keyboard['e'] = 1
	keyboard['r'] = 1
	keyboard['t'] = 1
	keyboard['y'] = 1
	keyboard['u'] = 1
	keyboard['i'] = 1
	keyboard['o'] = 1
	keyboard['p'] = 1

	keyboard['a'] = 2
	keyboard['s'] = 2
	keyboard['d'] = 2
	keyboard['f'] = 2
	keyboard['g'] = 2
	keyboard['h'] = 2
	keyboard['j'] = 2
	keyboard['k'] = 2
	keyboard['l'] = 2

	keyboard['z'] = 3
	keyboard['x'] = 3
	keyboard['c'] = 3
	keyboard['v'] = 3
	keyboard['b'] = 3
	keyboard['n'] = 3
	keyboard['m'] = 3

	count := 0
	for _, word := range words {
		row := keyboard[toLower(word[0])]
		for j := 1; j < len(word); j++ {
			if keyboard[toLower(word[j])] != row {
				row = 0
				break
			}
		}
		if row > 0 {
			words[count] = word
			count++
		}
	}

	return words[:count]
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}