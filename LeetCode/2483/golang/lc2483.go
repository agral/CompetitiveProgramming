package lc2483

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium
// Solved on: 2025-12-26
func bestClosingTime(customers string) int {
	n := len(customers)

	// From the challenge description:
	// For every hour when the shop is open and no customers come, the penalty increases by 1.
	// For every hour when the shop is closed and customers come, the penalty increases by 1.

	// 1. Calculate the prefix sum (from 0'th to c'th hour) of all the hours with no clients:
	prefixSumN := make([]int, n+1)
	for c := range n {
		if customers[c] == 'N' {
			prefixSumN[c+1] = prefixSumN[c] + 1
		} else {
			prefixSumN[c+1] = prefixSumN[c]
		}
	}

	// 2. Calculate the suffix sum (from c'th to the last hour) of all the hours with clients:
	suffixSumY := make([]int, n+1)
	for c := n - 1; c >= 0; c-- {
		if customers[c] == 'Y' {
			suffixSumY[c] = suffixSumY[c+1] + 1
		} else {
			suffixSumY[c] = suffixSumY[c+1]
		}
	}

	// 3. At any given hour, the penalty is calculated as prefixSumN[h] + suffixSumY[h].
	// Pick h such that h incurs the lowest possible penalty. In case of ties, pick lowest such h.
	ans := 0
	ansPenalty := prefixSumN[0] + suffixSumY[0]
	for h := 1; h <= n; h++ {
		hPenalty := prefixSumN[h] + suffixSumY[h]
		if hPenalty < ansPenalty {
			ans = h
			ansPenalty = hPenalty
		}
	}
	return ans
}
