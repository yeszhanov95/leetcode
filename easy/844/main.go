package main

func backspaceCompare(S string, T string) bool {
	stackS, stackT := make([]rune, 0, len(S)), make([]rune, 0, len(T))

	for _, r := range S {
		if r == '#' {
			if len(stackS) > 0 {
				stackS = stackS[:len(stackS)-1]
			}
			continue
		}
		stackS = append(stackS, r)
	}

	for _, r := range T {
		if r == '#' {
			if len(stackT) > 0 {
				stackT = stackT[:len(stackT)-1]
			}
			continue
		}
		stackT = append(stackT, r)
	}

	if len(stackS) != len(stackT) {
		return false
	}

	for i := range stackS {
		if stackS[i] != stackT[i] {
			return false
		}
	}

	return true
}