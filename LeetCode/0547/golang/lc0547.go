package lc0547

type UnionFind struct {
	count int
	id    []int
	rank  []int
}

func (uf *UnionFind) find(u int) int {
	if uf.id[u] == u {
		return u
	}
	uf.id[u] = uf.find(uf.id[u])
	return uf.id[u]
}

func (uf *UnionFind) unionByRank(lhs int, rhs int) {
	a := uf.find(lhs)
	b := uf.find(rhs)
	if a == b {
		return
	}

	if uf.rank[a] < uf.rank[b] {
		uf.id[a] = b
	} else if uf.rank[a] > uf.rank[b] {
		uf.id[b] = a
	} else { // uf.rank[a] == uf.rank[b]:
		uf.id[a] = b
		uf.rank[b]++
	}
	uf.count--
}

func makeUnionFind(size int) UnionFind {
	instance := UnionFind{}
	instance.count = size
	instance.id = make([]int, size)
	instance.rank = make([]int, size)
	for i := range size {
		instance.id[i] = i
	}
	return instance
}

// Runtime complexity: O(size^2 * log(size)), where size is the length of `isConnected`.
// Auxiliary space: O(size).
func findCircleNum(isConnected [][]int) int {
	size := len(isConnected)
	uf := makeUnionFind(size)
	for i := range size {
		for j := i; j < size; j++ {
			if isConnected[i][j] == 1 {
				uf.unionByRank(i, j)
			}
		}
	}
	return uf.count
}
