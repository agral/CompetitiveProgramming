package lc0841

// Runtime complexity: O(|V| + |E|),
// Auxiliary space: O(|V|),
// where |V|: total count of vertices; |E|: total count of edges,
// of the underlying graph of rooms (vertices) and keys (edges).
func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms)) // default-initialized to all false.

	var dfs func(roomNo int) // have to pre-declare this; it's a recursive function.
	dfs = func(roomNo int) {
		visited[roomNo] = true
		for _, roomKey := range rooms[roomNo] {
			if !visited[roomKey] {
				dfs(roomKey)
			}
		}
	}
	dfs(0)

	for _, isVisited := range visited {
		if !isVisited {
			return false
		}
	}
	return true
}
