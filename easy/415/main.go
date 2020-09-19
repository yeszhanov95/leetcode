package main

func addStrings(num1 string, num2 string) string {
	res := make([]byte, len(num1)+1)
	if len(num1) < len(num2) {
		res = make([]byte, len(num2)+1)
	}

	for carry, i, j, k := byte(0), len(num1)-1, len(num2)-1, len(res)-1; k >= 0; carry, i, j, k = carry/10, i-1, j-1, k-1 {
		if i >= 0 {
			carry += num1[i] - '0'
		}
		if j >= 0 {
			carry += num2[j] - '0'
		}
		res[k] = (carry % 10) + '0'
	}

	if res[0] == '0' {
		return string(res[1:])
	}
	return string(res)
}