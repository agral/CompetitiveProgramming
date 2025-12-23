package lc1870

import "math"

// Runtime complexity: O(n*log(1e7)) == O(n)
// Auxiliary space: O(1)
// Subjective level: medium
// Solved on: 2025-12-23
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)

	// as each train takes at least 1 full hour to complete its ride,
	// it is impossible to arrive on time if hour <= num_trains - 1:
	if hour <= float64(n-1) {
		return -1
	}

	canMakeOnTime := func(maxTime float64, trainSpeed float64) bool {
		travelTime := 0.0
		for d := 0; d < n-1; d++ {
			travelTime += math.Ceil(float64(dist[d]) / trainSpeed)
		}
		travelTime += float64(dist[n-1]) / trainSpeed
		return travelTime <= maxTime
	}

	low, high := 1, 10_000_000 // low and high boundaries for optimal train speed.
	for low < high {
		var mid int = low + (high-low)/2
		if canMakeOnTime(hour, float64(mid)) {
			// Midpoint speed makes the train reach target in time,
			// making it an upper bound for true optimal speed.
			high = mid
		} else {
			// Midpoint speed is insufficient to reach target in time,
			// making (midpoint+1) a lower bound for true optimal speed.
			low = mid + 1
		}
	}
	return low
}
