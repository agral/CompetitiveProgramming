package lc1401

// Runtime complexity: O(1)
// Auxiliary space: O(1)
func clamp(value int, minValue int, maxValue int) int {
	return max(minValue, min(maxValue, value))
}

// Runtime complexity: O(1)
// Auxiliary space: O(1)
// Subjective level: medium+
// Solved on: 2026-01-30
func checkOverlap(radius int, x int, y int, x1 int, y1 int, x2 int, y2 int) bool {
	// get the closest point from the circle's center that lies in the rectangle's bounds:
	closestX := clamp(x, x1, x2)
	closestY := clamp(y, y1, y2)

	// calculate the distance from closest point to the center of the circle:
	distX, distY := x-closestX, y-closestY

	// if this distance is <= the circle's radius, both figures are intersecting.
	// note: using distances-squared, as it is much more efficient.
	return (distX*distX)+(distY*distY) <= (radius * radius)
}
