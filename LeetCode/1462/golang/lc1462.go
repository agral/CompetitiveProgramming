package lc1462

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// Solved using Floyd-Warshall algorithm.
	sz := len(queries)
	ans := make([]bool, sz)

	// isPrerequisite[i][j] holds whether course i is a prerequisite of course j (i --> j on a dependency graph)
	isPrerequisite := make([][]bool, numCourses)
	for i := range numCourses {
		isPrerequisite[i] = make([]bool, numCourses)
	}
	for _, prq := range prerequisites {
		isPrerequisite[prq[0]][prq[1]] = true
	}

	for c := 0; c < numCourses; c++ {
		for u := 0; u < numCourses; u++ {
			for v := 0; v < numCourses; v++ {
				// WA #01
				//isPrerequisite[c][v] = isPrerequisite[c][v] || (isPrerequisite[c][u] && isPrerequisite[u][v])
				isPrerequisite[u][v] = isPrerequisite[u][v] || (isPrerequisite[u][c] && isPrerequisite[c][v])
			}
		}
	}

	for i, query := range queries {
		ans[i] = isPrerequisite[query[0]][query[1]]
	}

	return ans
}
