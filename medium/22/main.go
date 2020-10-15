package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	helper(0, 0, n, "", &res)
	return res
}

func helper(opening, closing, n int, str string, res *[]string) {
	if len(str) == n * 2 {
		*res = append(*res, str)
		return
	}
	if opening < n {
		helper(opening + 1, closing, n, str + "(", res)
	}
	if closing < opening {
		helper(opening, closing + 1, n, str + ")", res)
	}
}