package lc3296

import (
	"math"
	"slices"
)

// Runtime complexity:
//   - it's complicated :-)
//   - O(len(workerTimes)*log(mountainHeight^2) * (some further analysis needed!).
//   - surely less than O(len(workerTimes) * mountainHeight^2), and
//     surely more than O(len(workerTimes) * log(mountainHeight^2)).
//
// Auxiliary space: O(1).
// Subjective level: medium+. Tedious to analyze, rather easy to write
// (but Golang hates math, see conversions between int, int64 and float64 issued everywhere).
// This would have been a medium- in Python...
// Solved on: 2026-03-13
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	minWorkerTime := slices.Min(workerTimes)
	low := int64(1)
	high := int64(minWorkerTime) * int64(mountainHeight) * int64(mountainHeight+1) / int64(2)

	getHeightReducedInTime := func(seconds int64) int {
		ans := 0
		// the formula for height h removed by a worker who works
		// for `s` seconds and has `t` workerTime:
		// t * (1 + 2 + ... + h) <= s  (note: we're bound to whole ints, so <= and not ==)
		// t * h * (h+1) / 2 <= s
		// (h^2 + h) / 2 <= s/t
		// h^2 + h <= 2s/t
		// h^2 <= 2s/t - h // div-by-h:
		// h <= (-1 + (sqrt(1 + 8s/t))) / 2
		for _, time := range workerTimes {
			ans += int(math.Sqrt(8*float64(seconds)/float64(time)+1)-1) / 2
		}
		return ans
	}

	for low < high {
		mid := (low + high) / 2
		if getHeightReducedInTime(mid) < mountainHeight {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return int64(low)
}
