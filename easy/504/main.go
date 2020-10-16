package main

func convertToBase7(num int) string {
	if num == 0 { return "0" }
	if num < 0 { return "-" + convertToBase7(-num) }
	slice := make([]byte, 0)
	for num > 0 {
		slice = append(slice, byte(num % 7) + '0')
		num /= 7
	}
	for i, j := 0, len(slice) - 1; i < j; i, j = i + 1, j - 1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return string(slice)
}