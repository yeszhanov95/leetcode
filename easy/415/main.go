package main

func addStrings(num1 string, num2 string) string {
	zeroVal := []byte("0")[0]
	cacheNums, cacheChars := make(map[byte]int, 10), make(map[int]byte, 10)
	for i := 0; i < 10; i++ {
		c := zeroVal + byte(i)
		cacheNums[c] = i
		cacheChars[i] = c
	}

	result := make([]byte, len(num1) + len(num2))
	idx, i, j, remainder := len(result) - 1, len(num1)-1, len(num2)-1, 0

	for ; i >= 0 && j >= 0; i, j, idx = i-1, j-1, idx-1 {
		num := cacheNums[num1[i]] + cacheNums[num2[j]] + remainder
		result[idx] = cacheChars[num % 10]
		if num >= 10 {
			remainder = 1
		} else {
			remainder = 0
		}
	}

	if i < 0 {
		for ; j >= 0; j, idx = j-1, idx-1 {
			num := cacheNums[num2[j]] + remainder
			result[idx] = cacheChars[num % 10]
			if num >= 10 {
				remainder = 1
			} else {
				remainder = 0
			}
		}
	} else if j < 0 {
		for ; i >= 0; i, idx = i-1, idx-1 {
			num := cacheNums[num1[i]] + remainder
			result[idx] = cacheChars[num % 10]
			if num >= 10 {
				remainder = 1
			} else {
				remainder = 0
			}
		}
	}

	if remainder > 0 {
		result[idx] = cacheChars[1]
		idx--
	}

	return string(result[idx+1:])
}