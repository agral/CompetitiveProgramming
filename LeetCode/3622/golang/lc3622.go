package lc3622

// Runtime complexity: O(log10(n))
// Auxiliary space: O(1)
// Subjective level: trivially easy.
// Solved on: 2026-08-22
func checkDivisibility(n int) bool {
	digitSum := 0
	digitProduct := 1

	k := n
	for k > 0 {
		d := k % 10
		digitSum += d
		digitProduct *= d
		k /= 10
	}

	return n%(digitSum+digitProduct) == 0
}
