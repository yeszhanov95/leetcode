package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjMatrix := make(map[int][]int)
	finished, inProgress := make(map[int]bool), make(map[int]bool)

	for _, prereq := range prerequisites {
		want := prereq[0]
		need := prereq[1]
		adjMatrix[want] = append(adjMatrix[want], need)
	}

	for i := 0; i < numCourses; i++ {
		if finished[i] {
			continue
		}
		if hasCycle(i, adjMatrix, finished, inProgress) {
			return false
		}
	}

	return true
}

func hasCycle(course int, adjMatrix map[int][]int, finished, inProgress map[int]bool) bool {
	if inProgress[course] {
		return true
	}

	needs, ok := adjMatrix[course]
	if !ok {
		finished[course] = true
		return false
	}

	inProgress[course] = true
	for _, need := range needs {
		if finished[need] {
			continue
		}
		if hasCycle(need, adjMatrix, finished, inProgress) {
			return true
		}
	}

	inProgress[course] = false
	finished[course] = true

	return false
}