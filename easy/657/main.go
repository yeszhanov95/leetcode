package main

func judgeCircle(moves string) bool {
	x, y := 0, 0
	runes := []rune(moves)
	for i := 0; i < len(runes); i++ {
		move := runes[i]
		switch move {
		case 'R': x += 1
		case 'D': y -= 1
		case 'L': x -= 1
		case 'U': y += 1
		}
	}
	return x == 0 && y == 0
}