package lc2147

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: medium+
// Solved on: 2025-12-14
func numberOfWays(corridor string) int {
	const MOD64 int64 = 1_000_000_007
	lastSeatIdx := -1
	totalSeats := 0
	ans := int64(1)

	for i, elem := range corridor {
		if elem == 'S' {
			totalSeats += 1
			if totalSeats%2 == 1 && totalSeats >= 3 {
				ans = (ans * int64(i-lastSeatIdx)) % MOD64
			}
			lastSeatIdx = i
		}
	}

	// Corner case #01: if there are no seats at all, there can be no segments at all.
	// Corner case #02: if totalSeats is even, there is no way to put the odd one to use.
	// both these values are known only at the end of a full scan of the array.
	if totalSeats == 0 || totalSeats%2 == 1 {
		return 0
	}

	return int(ans)
}
