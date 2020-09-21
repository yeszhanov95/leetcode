package main

func numberOfSubstrings(s string) int {
	counter := make([]int, 3)
	var res, tmp int
	for i, j := 0, 0; i < len(s); i++ {
		counter[int(s[i]-'a')]++
		for counter[0] >= 1 && counter[1] >= 1  && counter[2] >= 1  {
			counter[s[j]-'a']--
			tmp, j = tmp + 1, j + 1
		}
		res += tmp
	}
	return res
}