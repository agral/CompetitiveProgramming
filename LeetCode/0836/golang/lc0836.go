package lc0836

// Runtime complexity: O(1)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-09-14
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return rec1[0] < rec2[2] &&
		rec1[2] > rec2[0] &&
		rec1[1] < rec2[3] &&
		rec1[3] > rec2[1]
}
