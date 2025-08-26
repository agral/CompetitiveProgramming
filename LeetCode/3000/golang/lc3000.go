package lc3000

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
func areaOfMaxDiagonal(dimensions [][]int) int {
	maxDiagonalSquared := 0
	maxArea := 0
	for i := range dimensions {
		diagonalSquared := (dimensions[i][0] * dimensions[i][0]) + (dimensions[i][1] * dimensions[i][1])
		if diagonalSquared > maxDiagonalSquared {
			maxDiagonalSquared = diagonalSquared
			maxArea = dimensions[i][0] * dimensions[i][1]
		} else if diagonalSquared == maxDiagonalSquared {
			maxArea = max(maxArea, dimensions[i][0]*dimensions[i][1])
		}
	}
	return maxArea
}
