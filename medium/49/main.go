package main

func groupAnagrams(strs []string) [][]string {
	var numOfGroups int
	result := make([][]string, len(strs))
	cache := make(map[[26]byte]int, len(strs))

	for _, str := range strs {
		bytes := [26]byte{}
		for _, v := range str {
			bytes[v-'a']++
		}

		if _, ok := cache[bytes]; ok {
			result[cache[bytes]] = append(result[cache[bytes]], str)
		} else {
			result[numOfGroups] = []string{str}

			cache[bytes] = numOfGroups
			numOfGroups++
		}
	}

	return result[:numOfGroups]
}