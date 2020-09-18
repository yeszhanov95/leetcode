package main

func repeatedSubstringPattern(s string) bool {
	divider := len(s) / 2
	var left1, right1, left2, right2 int

	for ; divider > 0; divider-- {
		if len(s) % divider != 0 {
			continue
		}
		res := true

		left1, right1, left2, right2 = 0, divider, divider, divider*2
		for left2 < len(s) {
			if s[left1:right1] != s[left2:right2] {
				res = false
				break
			}
			left1, right1, left2, right2 = left2, right2, left2+divider, right2+divider
		}

		if res == true {
			return true
		}
	}

	return false
}