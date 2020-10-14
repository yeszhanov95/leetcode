package main

func combine(n int, k int) [][]int {
	if n == k {
		cur := make([]int, k)
		for i := 1; i <= n; i++ {
			cur[i-1] = i
		}
		return [][]int{cur}
	}

	res, cur := make([][]int, 0), make([]int, 0, k)
	helper(1, n, cur, &res)
	return res
}

func helper(start, limit int, cur []int, res *[][]int) {
	if len(cur) == cap(cur) {
		*res = append(*res, append([]int{}, cur...))
		return
	}
	if len(cur) + (limit - start + 1) < cap(cur) {
		return
	}

	for i := start; i <= limit; i++ {
		cur = append(cur, i)
		helper(i + 1, limit, cur, res)
		cur = cur[:len(cur)-1]
	}
}