package main

func checkIfExist(arr []int) bool {
	for i, v := range arr {
		for i2, v2 := range arr {
			if i == i2 {
				continue
			}
			if (v2 % 2 == 0 && v2 / 2 == v) || v2 * 2 == v {
				return true
			}
		}
	}
	return false
}