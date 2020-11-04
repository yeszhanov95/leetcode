package main

func maxAreaOfIsland(grid [][]int) int {
	var maxArea, curArea int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 { continue }
			curArea = 0
			dfs(i, j, &grid, &curArea)
			if curArea > maxArea { maxArea = curArea }
		}
	}
	return maxArea
}

func dfs(i, j int, grid *[][]int, area *int) {
	if i < 0 || j < 0 || i == len(*grid) || j == len((*grid)[0]) || (*grid)[i][j] == 0 {
		return
	}
	*area++
	(*grid)[i][j] = 0
	dfs(i + 1, j, grid, area)
	dfs(i - 1, j, grid, area)
	dfs(i, j + 1, grid, area)
	dfs(i, j - 1, grid, area)
}