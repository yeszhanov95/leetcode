package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjLists, order := make(map[int][]int), make([]int, numCourses)
	for _, prereq := range prerequisites {
		adjLists[prereq[1]] = append(adjLists[prereq[1]], prereq[0])
	}

	visited, recStack := make(map[int]bool, numCourses), make(map[int]bool, numCourses)
	for course, i := 0, numCourses - 1; course < numCourses; course++ {
		if visited[course] { continue }
		i = dfsTopSort(course, i, adjLists, visited, recStack, order)
		if i == -2 { return []int{} }
	}

	return order
}

func dfsTopSort(course, i int, adjLists map[int][]int, visited, recStack map[int]bool, order []int) int {
	if recStack[course] { return -2 }
	if visited[course] { return i }

	visited[course], recStack[course] = true, true
	if _, ok := adjLists[course]; ok {
		for _, neighbour := range adjLists[course] {
			i = dfsTopSort(neighbour, i, adjLists, visited, recStack, order)
			if i == -2 { return -2 }
		}
	}

	recStack[course] = false
	order[i] = course
	return i - 1

}