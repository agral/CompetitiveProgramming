package lc1266

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-01-12
func minTimeToVisitAllPoints(points [][]int) int {
	n := len(points)
	ans := 0
	for i := 1; i < n; i++ {
		ans += max(abs(points[i][0]-points[i-1][0]), abs(points[i][1]-points[i-1][1]))
	}
	return ans
}
