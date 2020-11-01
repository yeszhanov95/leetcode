package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if newColor != image[sr][sc] {
		dfs(sr, sc, image[sr][sc], newColor, image)
	}
	return image
}

func dfs(i, j, prevColor, newColor int, image [][]int) {
	if i < 0 || i == len(image) || j < 0 || j == len(image[0]) || image[i][j] != prevColor {
		return
	}
	image[i][j] = newColor
	dfs(i, j + 1, prevColor, newColor, image)
	dfs(i, j - 1, prevColor, newColor, image)
	dfs(i + 1, j, prevColor, newColor, image)
	dfs(i - 1, j, prevColor, newColor, image)
}