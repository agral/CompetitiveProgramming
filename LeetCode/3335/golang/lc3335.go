package lc3335

// Recursive. Overflows the stack quite fast.
// Does not even do the mod, as the stack is overflown first :-)
// But used this to test the actual solution for correctness.
func naiveGetLength(offset int, t int) int {
	if t == 0 {
		return 1
	}

	if offset == 25 {
		return naiveGetLength(0, t-1) + naiveGetLength(1, t-1)
	}

	return naiveGetLength(offset+1, t-1)
}

// Runtime complexity: O(len(s) + 26*t); usually t dominates.
// Auxiliary space: O(26 + 26) == O(1).
func lengthAfterTransformations(s string, t int) int {
	const MOD = 1_000_000_007
	bytestring := []byte(s)
	lettersCount := make([]int, 26)
	for _, ch := range bytestring {
		lettersCount[int(ch-'a')]++
	}

	for range t {
		newCount := make([]int, 26)
		// Handles a->b, b->c, ..., y->z; but not z->ab
		for letterOffset := range 25 {
			newCount[letterOffset+1] = lettersCount[letterOffset]
		}
		// handles z->ab:
		newCount[0] = lettersCount[25]
		newCount[1] = (newCount[1] + lettersCount[25]) % MOD // the only one that can become greater is "b".
		lettersCount = newCount
	}
	ans := 0
	for _, letterCount := range lettersCount {
		ans = (ans + letterCount) % MOD
	}
	return ans
}
