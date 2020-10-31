package main

func numIslands(grid [][]byte) int {
	var islandsNumber int
	for line := range grid {
		for val := range grid[line] {
			if grid[line][val] == '1' {
				islandsNumber++
				dfs(line, val, grid)
			}
		}
	}
	return islandsNumber
}

func dfs(i, j int, grid [][]byte) {
	grid[i][j] = '0'
	if j + 1 < len(grid[0]) && grid[i][j+1] == '1' {
		dfs(i, j + 1, grid)
	}
	if j - 1 >= 0 && grid[i][j-1] == '1' {
		dfs(i, j - 1, grid)
	}
	if i + 1 < len(grid) && grid[i+1][j] == '1' {
		dfs(i + 1, j, grid)
	}
	if i - 1 >= 0 && grid[i-1][j] == '1' {
		dfs(i - 1, j, grid)
	}
}