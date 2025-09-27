package lc0812

import "math"

// Returns the area of the triangle formed on three planar (2D) points p1, p2, p3.
// Uses the Shoelace's formula to calculate the result.
func getTriangleArea(p1 []int, p2 []int, p3 []int) float64 {
	return 0.5 * math.Abs(float64(p1[0]*(p2[1]-p3[1])+p2[0]*(p3[1]-p1[1])+p3[0]*(p1[1]-p2[1])))
}

// Runtime complexity: O(n^3)
// Auxiliary space: O(1)
// Subjective level: easy, but had to search up the Shoelace's formula
func largestTriangleArea(points [][]int) float64 {
	ans := 0.0
	// I don't think there's an easy way to skip brute-forcing all possible triplets of points.
	LEN_POINTS := len(points)
	for first := 0; first < LEN_POINTS-2; first++ {
		for second := first + 1; second < LEN_POINTS-1; second++ {
			for third := second + 1; third < LEN_POINTS; third++ {
				ans = math.Max(ans, getTriangleArea(points[first], points[second], points[third]))
			}
		}
	}
	return ans
}
