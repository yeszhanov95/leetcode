package main

var matrix [][]int
var visited []bool

func findCircleNum(M [][]int) int {
	var circles int
	matrix, visited = M, make([]bool, len(M))
	for student := 0; student < len(matrix); student++ {
		if !visited[student] {
			circles++
			dfs(student)
		}
	}
	return circles
}

func dfs(student int) {
	if visited[student] { return }
	visited[student] = true
	for otherStudent := 0; otherStudent < len(matrix); otherStudent++ {
		if matrix[student][otherStudent] == 1 && student != otherStudent {
			dfs(otherStudent)
		}
	}
}