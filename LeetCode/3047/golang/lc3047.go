package lc3047

// Runtime complexity: O(L^2); where L is the length of bottomLeft
// (which is equal to the length of topRight). Every such rectangle has to be
// compared pairwise to all the other rectangles.
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2026-01-17
func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	minSquareSide := 0
	for i := 0; i < len(bottomLeft); i++ {
		for j := i + 1; j < len(bottomLeft); j++ {
			intersectionX := min(topRight[i][0], topRight[j][0]) - max(bottomLeft[i][0], bottomLeft[j][0])
			intersectionY := min(topRight[i][1], topRight[j][1]) - max(bottomLeft[i][1], bottomLeft[j][1])
			minSquareSide = max(minSquareSide, min(intersectionX, intersectionY))
		}
	}
	return int64(minSquareSide) * int64(minSquareSide)
}
