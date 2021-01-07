package main

func dailyTemperatures(T []int) []int {
	n := len(T)
	if n == 1 { return []int{0} }

	stack := make([]int, 0)
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 {
			if T[stack[len(stack)-1]] > T[i] {
				result[i] = stack[len(stack)-1] - i
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, i)
	}

	return result
}