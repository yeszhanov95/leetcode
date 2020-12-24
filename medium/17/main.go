package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	dict := map[rune]string {
		'2' : "abc", '3' : "def", '4' : "ghi",
		'5' : "jkl", '6' : "mno", '7' : "pqrs",
		'8' : "tuv", '9' : "wxyz",
	}
	combinations := make([]string, 0)
	for _, r := range digits {
		combinations = append(combinations, dict[r])
	}

	result := make([]string, 0)
	helper(&combinations, &result, 0, []byte{})

	return result
}

func helper(combinations, result *[]string, pos int, curr []byte) {
	if pos == len(*combinations) {
		*result = append(*result, string(curr))
		return
	}

	combination := (*combinations)[pos]
	for i := range combination {
		helper(combinations, result, pos + 1, append(curr, combination[i]))
	}
}