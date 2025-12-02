package lc3623

// Runtime complexity: O(n)
// Auxiliary space: O(n) // in the worst case; close to O(log(n)) for the average case.
// Subjective level: medium.
// Solved on: 2025-12-02
func countTrapezoids(points [][]int) int {
	const MOD64 int64 = 1_000_000_007

	// approach: count the number of distinct points that have the same y-coordinate.
	// but given that all the points are pairwise different, their X-coordinate does not even
	// need to get stored (it's different from all the other X-coordinates on this line, that's
	// what matters). All that needs to get stored is the count of points with same Y-coordinate.
	// These points can form the base of a trapezoid, as long as there's at least two of them.
	ys := make(map[int]int64)

	for _, point := range points {
		ys[point[1]] += 1
	}

	// when there are n points with same y-coordinate, a base can be created in (n)*(n-1)/2 ways.
	// when there are many lines, simply multiply the number of combinations so far with this line's
	// count of ways to create a base.
	ans := int64(0) // note: will return 0 unless there's at least two distinct y-coords with at least two points each.
	total := int64(0)
	//numBaseLines := 0
	for _, y := range ys {
		// ignore the lines that have a singular point - these can not create a base.
		if y >= 2 {
			//numBaseLines += 1
			numLines := y * (y - 1) / 2
			//note: by switching up the order ans and total are counted,
			// it's no longer necessary to account for numBaseLines -
			// the first line counts against total, but does not change the ans.
			ans += total * numLines
			ans %= MOD64
			total += numLines
			total %= MOD64
		}
	}

	// if numBaseLines < 2 {
	// 	return 0
	// }

	return int(ans)
}
