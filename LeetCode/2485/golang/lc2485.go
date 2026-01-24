package lc2485

func sumNums(start int, stop int) int {
	mid := 0.5 * (float64(start) + float64(stop))
	return int(float64(stop-start+1) * mid)
}

// Runtime complexity: O(log(1000)) == O(1).
// Auxiliary space: O(1).
// Subjective level: medium.
// Solved on: 2026-01-24
// Note: turns out that there's just 4 values of n in 1-1000 range that have a pivot number.
// 1, 8, 49, and 288. So probably lots of solutions out there cheese it.
func pivotInteger(n int) int {
	// binary search for the pivot integer.
	l := 1
	h := n
	for l <= h {
		mid := (l + h) / 2
		sumOneMid := sumNums(1, mid)
		sumMidN := sumNums(mid, n)
		if sumOneMid == sumMidN {
			// found the pivot integer!
			return mid
		} else {
			if sumOneMid < sumMidN {
				l = mid + 1
			} else {
				h = mid - 1
			}
		}
	}
	return -1
}
