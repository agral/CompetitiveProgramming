package lc3607

type Set map[int]bool

// Runtime complexity:
// Auxiliary space:
// Subjective level:
func processQueries(c int, connections [][]int, queries [][]int) []int {
	roots := make([]int, c+1) //c+1 for the stations are 1-indexed, not 0-indexed.
	for i := 1; i <= c; i++ {
		roots[i] = i
	}

	// recursive func, needs to be declared and only then defined.
	var find func(i int) int
	find = func(i int) int {
		if roots[i] == i {
			return i
		}
		roots[i] = find(roots[i])
		return roots[i]
	}

	union := func(a int, b int) {
		rootA := find(a)
		rootB := find(b)
		if rootA != rootB {
			if rootA < rootB {
				roots[rootB] = rootA
			} else {
				roots[rootA] = rootB
			}
		}
	}

	for _, c := range connections {
		union(c[0], c[1])
	}

	gridMembers := make([][]int, c+1)
	for i := 1; i <= c; i++ {
		gridMembers[i] := []int{}
	}
	for i := 1; i <= c; i++ {
		root := find(i)
		if
	}

	ans := []int{}
	return ans
}
