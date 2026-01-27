package lc3650

import (
	"container/heap"
	"math"
)

type DistNode struct {
	distance int
	index    int
}

type PriorityQueue []DistNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(a, b int) bool {
	return pq[a].distance < pq[b].distance
}

func (pq PriorityQueue) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(DistNode))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	l := len(old)
	item := old[l-1]
	*pq = old[0 : l-1]
	return item
}

// Runtime complexity: O(|E+V|*log|V|), probably. Don't know container/heap's internals, but surely this.
// Auxiliary space: O(|V|)
// Subjective level: hard. container/heap is still subjectively hard to use.
// Solved on: 2026-01-27
func minCost(n int, edges [][]int) int {
	// Note: it's guaranteed that weight values are never negative (1 <= wi <= 1000),
	// so it's enough to add reversed edges to the graph and run Dijkstra's algorithm.

	// 1. Process the edges. Construct the graph, but for each edge
	// also add a reversed edge with double the cost.
	g := make([][][2]int, n)
	for _, edge := range edges {
		from, to, cost := edge[0], edge[1], edge[2]
		g[from] = append(g[from], [2]int{to, cost})
		g[to] = append(g[to], [2]int{from, 2 * cost})
	}

	// 2. Mark each node as quite distant; except the 0th node, which is immediately reachable.
	bigDistance := math.MaxInt / 2 // half of MaxInt, just to be sure that no overflow occurs.
	dist := make([]int, n)
	for i := 1; i < n; i++ {
		dist[i] = bigDistance
	}
	// implicitly: dist[0] has been zeroed when creating with `make`, so it remains that dist[0] == 0.

	// 3. Run Dijkstra's algorithm to find the distance between nodes 0 and n-1.
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, DistNode{0, 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(DistNode)
		// don't process paths with distances larger than currently known optimal:
		if curr.distance <= dist[curr.index] {
			// when node n-1 is reached for the first time, it's already an optimal path (dijkstra's cool property).
			if curr.index == n-1 {
				return curr.distance
			}
			for _, from := range g[curr.index] {
				to, cost := from[0], from[1]
				newDistance := curr.distance + cost
				if newDistance < dist[to] {
					dist[to] = newDistance
					heap.Push(pq, DistNode{newDistance, to})
				}
			}
		}
	}

	// Node n-1 is not reachable at all from node 0.
	return -1
}
