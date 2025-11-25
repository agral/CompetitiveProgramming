package lc1015

// Runtime complexity: O(k)
// Auxiliary space: O(k)
// Subjective level: medium/hard
// Solved on: 2025-11-25
func smallestRepunitDivByK(k int) int {
	// A string of '1's will never result in an even number.
	if k%2 == 0 {
		return -1
	}

	rem := 0 // the number n gets large pretty fast; but it's sufficient to store the remainder.
	seen := map[int]bool{}
	for length := 1; length <= k; length++ {
		rem = (10*rem + 1) % k
		if rem == 0 {
			return length
		}
		if seen[rem] {
			// a loop has been detected, adding more '1's will only make the cycle repeat itself again.
			return -1
		}
		seen[rem] = true
	}

	return -1
}
