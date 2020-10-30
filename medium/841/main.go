package main

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	dfs(0, visited, rooms)
	for i := range visited {
		if !visited[i] { return false }
	}
	return true
}

func dfs(room int, visited []bool, rooms [][]int) {
	visited[room] = true
	for _, accessibleRoom := range rooms[room] {
		if !visited[accessibleRoom] {
			dfs(accessibleRoom, visited, rooms)
		}
	}
} 