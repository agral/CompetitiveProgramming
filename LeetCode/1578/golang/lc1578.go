package lc1578

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
func minCost(colors string, neededTime []int) int {
	previousTime := neededTime[0]
	ans := 0
	for i := 1; i < len(colors); i++ {
		if colors[i-1] == colors[i] {
			// The time for cutting the faster of the two is added to the immediate answer.
			// The time for cutting the slower of the two remains to be considered for the
			// next cutting, if the next balloon is also the same color.

			// don't want to implement max/min for ints. Go has them only for float64 for some reason.
			if previousTime < neededTime[i] {
				ans += previousTime
				previousTime = neededTime[i]
			} else {
				ans += neededTime[i]
				// previousTime remains as it was.
			}
		} else {
			// Don't cut anything, but remember this balloon's time in case the next is also same color.
			previousTime = neededTime[i]
		}
	}
	return ans
}
