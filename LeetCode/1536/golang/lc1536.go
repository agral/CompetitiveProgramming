package lc1536

// Runtime complexity: O(n^2)
// Auxiliary space: O(n)
// Subjective level: medium
// Solved on: 2026-03-02
func minSwaps(grid [][]int) int {
	L := len(grid)
	ans := 0

	// Count the number of suffix zeroes for each row:
	suffixZeroes := make([]int, L)
	for numRow := range L {
		var i = 0
		for i = L - 1; i >= 0 && grid[numRow][i] == 0; i-- {
		}
		suffixZeroes[numRow] = L - 1 - i
	}

	const NO_SUCH_ROW = -1
	for i := range L {
		minimumZeroes := L - 1 - i
		// Find the first row that has suffixZeroes >= minimumZeroes,
		// in [i..n]. Note that [0..i-1] is already used and can not be reused.
		matchingRowIdx := NO_SUCH_ROW
		for k := i; matchingRowIdx == -1 && k < L; k++ {
			if suffixZeroes[k] >= minimumZeroes {
				matchingRowIdx = k
			}
		}
		if matchingRowIdx == NO_SUCH_ROW {
			return NO_SUCH_ROW
		}
		// now matchingRowIdx contains the index of first row where suffixZeroes >= minimumZeroes.
		// Move grid[matchingRowIdx] to grid[i]:
		for r := matchingRowIdx; r > i; r-- {
			suffixZeroes[r] = suffixZeroes[r-1]
		} // in the end, "forget" about suffixZeroes[i]. That's used up now.
		ans += matchingRowIdx - i
	}
	return ans
}
