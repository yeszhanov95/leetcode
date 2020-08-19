package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	needleArr := []byte(needle)
	haystackArr := []byte(haystack)

	for i := 0; i + len(needleArr) <= len(haystackArr); i++ {
		if haystackArr[i] == needleArr[0] && compare(haystackArr[i : i + len(needleArr)], needleArr) {
			return i
		}
	}

	return -1
}

func compare(first, second []byte) bool {
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}