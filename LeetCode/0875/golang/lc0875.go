package lc0875

import "slices"

// Runtime complexity: O(n * log(max(piles)))
// Auxiliary space: O(1)
func minEatingSpeed(piles []int, h int) int {
	low := 1
	high := slices.Max(piles)
	for low < high {
		mid := (low + high) / 2

		// Calculate the number of hours necessary to eat all the piles at speed `mid` bananas per hour:
		hours := 0
		for _, pile := range piles {
			hours += 1 + (pile-1)/mid // === math.Ceil((float64 pile / float64 mid), but a bit faster
		}
		if hours > h {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
