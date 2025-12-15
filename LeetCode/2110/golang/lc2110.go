package lc2110

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2025-12-15
func getDescentPeriods(prices []int) int64 {
	ans := int64(1)
	currSmooth := 0
	for idx := 1; idx < len(prices); idx++ {
		if prices[idx-1]-prices[idx] == 1 {
			currSmooth += 1
		} else {
			currSmooth = 0
		}
		ans += int64(currSmooth + 1)
	}
	return ans
}
