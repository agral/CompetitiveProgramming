package lc1975

import "math"

func abs(val int) int64 {
	if val < 0 {
		return int64(-val)
	}
	return int64(val)
}

// Runtime complexity: O(H*W), where H, W are matrix dimensions. This is optimal.
// Auxiliary space: O(1), optimal.
// Subjective level: easy/medium
// Solved on: 2026-01-05
func maxMatrixSum(matrix [][]int) int64 {
	// Approach: using the flip-sign operation any number of times,
	// we can flip the sign of any pair of entries. Depending on whether the total
	// count of initially negative entries is even or odd, there will be no negatives (even),
	// or one negative value remaining in the end.
	numNegatives := 0
	minAbsValue := int64(math.MaxInt64) // for some reason, I have to specify that MaxInt64 is an int64...
	absTotal := int64(0)
	for _, row := range matrix {
		for _, val := range row {
			absVal := abs(val)
			absTotal += absVal
			if absVal < minAbsValue {
				minAbsValue = absVal
			}
			if val < 0 {
				numNegatives += 1
			}
		}
	}
	if numNegatives%2 == 1 {
		// If there is one negative value remaining in the end, it needs to be subtracted twice from
		// the calculated absTotal: once because it was initially added, and once for the actual subtraction.
		absTotal -= 2 * int64(minAbsValue)
	}
	return absTotal
}
