package lc3531

import (
	"math"
	"slices"
)

// Runtime complexity: O(len(buildings))
// Auxiliary space: O(n) // where n == the size of the map, passed as first parameter.
// Subjective level: medium
// Solved on: 2025-12-11
func countCoveredBuildings(n int, buildings [][]int) int {
	ans := 0
	furthest_north := slices.Repeat([]int{math.MaxInt32}, n+1)
	furthest_east := make([]int, n+1)
	furthest_south := make([]int, n+1)
	furthest_west := slices.Repeat([]int{math.MaxInt32}, n+1)

	for _, coords := range buildings {
		x := coords[0]
		y := coords[1]
		furthest_north[x] = min(furthest_north[x], y)
		furthest_south[x] = max(furthest_south[x], y)
		furthest_east[y] = max(furthest_east[y], x)
		furthest_west[y] = min(furthest_west[y], x)
	}

	for _, coords := range buildings {
		x := coords[0]
		y := coords[1]
		if furthest_north[x] < y && y < furthest_south[x] &&
			furthest_west[y] < x && x < furthest_east[y] {
			ans += 1
		}
	}

	return ans
}
