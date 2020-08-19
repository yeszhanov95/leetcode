package main

func numberOfSteps (num int) int {
	res := 0

	if num == 0 {
		return res
	}

	for {
		if num == 0 {
			break
		}
		if num % 2 == 0 {
			num = num / 2
		} else {
			num--
		}
		res++
	}

	return res
}