package lc0149

type Point struct {
	x int
	y int
}

// Runtime complexity: O(log(max(x, y)))
// Auxiliary space: O(1)
func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// Runtime complexity: O(1)+O(gcd()) == O(log(max(x, y)))
// Auxiliary space: O(1)
func getSlope(a Point, b Point) Point {
	dx, dy := a.x-b.x, a.y-b.y
	if dy == 0 {
		return Point{a.y, 0}
	}
	if dx == 0 {
		return Point{0, a.x}
	}
	d := gcd(dx, dy)
	return Point{dx / d, dy / d}
}

// Runtime complexity: O(n^2)*O(getSlope), which is still O(n^2) in the end.
// Auxiliary space: O(n)
// Subjective level: hard
// Solved on: 2026-01-30
func maxPoints(points [][]int) int {
	ans := 0
	for i, point := range points {
		slopeToCount := map[Point]int{}
		numRepeatedPoints := 1
		numSameSlopePoints := 0
		p1 := Point{point[0], point[1]}
		for j := i + 1; j < len(points); j++ {
			point2 := points[j]
			p2 := Point{point2[0], point2[1]}
			if point[0] == point2[0] && point[1] == point2[1] {
				numRepeatedPoints += 1
			} else {
				slope := getSlope(p1, p2)
				_, isPresent := slopeToCount[slope]
				if isPresent {
					slopeToCount[slope] += 1
				} else {
					slopeToCount[slope] = 1
				}

				numSameSlopePoints = max(numSameSlopePoints, slopeToCount[slope])
			}
		}

		ans = max(ans, numRepeatedPoints+numSameSlopePoints)
	}
	return ans
}
