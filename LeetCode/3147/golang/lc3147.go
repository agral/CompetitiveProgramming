package lc3147

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: Medium/easy.
func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	dp := make([]int, n)
	copy(dp, energy)
	// note: built-in copy(target, source) lists target FIRST, source SECOND.
	// that looks unnatural (as one would e.g. `cp /src/file /out/file` in a shell);
	// golang has the args in inverted order for some reason.

	for i := n - 1 - k; i >= 0; i-- {
		dp[i] += dp[i+k]
	}
	idxMaxValue := 0
	for i := 1; i < n; i++ {
		if dp[i] > dp[idxMaxValue] {
			idxMaxValue = i
		}
	}
	return dp[idxMaxValue]
}
