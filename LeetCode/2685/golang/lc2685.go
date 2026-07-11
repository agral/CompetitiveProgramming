package lc2685

// Runtime complexity: O(|V| + |E|)
// Auxiliary space: O(|V| + |E|)
// Subjective level: medium+
// Solved on: 2026-07-11
func countCompleteComponents(n int, edges [][]int) int {
	graph := make([][]int, n)
	isVisited := make([]bool, n) // initialized to false by default, good.

	var dfs func(v int) (int, int)
	dfs = func(v int) (int, int) {
		isVisited[v] = true
		numNodes := 1
		numEdges := len(graph[v])
		for _, u := range graph[v] {
			if !isVisited[u] {
				n, e := dfs(u)
				numNodes += n
				numEdges += e
			}
		}
		return numNodes, numEdges
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	ans := 0
	for v, visited := range isVisited {
		if !visited {
			n, e := dfs(v)
			if (n-1)*n == e { // in other words, if this connected subgraph is a clique...
				ans += 1
			}
		}
	}
	return ans
}
