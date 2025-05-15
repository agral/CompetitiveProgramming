package lc1466

func abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

// Runtime complexity: O(n)
// Auxiliary space: O(2n) == O(n)
func minReorder(n int, connections [][]int) int {
	graph := make([][]int, n)

	for _, c := range connections {
		src := c[0]
		dest := c[1]
		graph[src] = append(graph[src], dest)
		graph[dest] = append(graph[dest], -src)
	}

	var dfs func(node int, previous int) int
	dfs = func(node int, previous int) int {
		change := 0
		for _, dest := range graph[node] {
			if previous != abs(dest) {
				if dest > 0 {
					change++
				}
				change += dfs(abs(dest), node)
			}
		}
		return change
	}
	return dfs(0, -1)
}
