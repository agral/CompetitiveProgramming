package lc1563

// Runtime complexity: O(n^3)
// Auxiliary space: O(n^2) (all the possible sub-sums set has n(n-1)/2 size,
// which is O(n^2) in O notation)
// Subjective level: hard
// Solved on: 2026-08-17
func stoneGameV(stoneValue []int) int {
	L := len(stoneValue)

	// prefixSum[i] holds the sum of values of stones up to but not including the ith.
	prefixSum := make([]int, L+1)
	for i := range L { // zero-initialized prefixSum[0] stays 0, that's OK.
		prefixSum[i+1] = prefixSum[i] + stoneValue[i]
	}

	memo := make([][]int, L)
	for row := range L {
		memo[row] = make([]int, L) // seems OK to keep it 0-initialized.
	}

	var dp func(left int, right int) int
	dp = func(left int, right int) int {
		if left == right { // note: not needed to do the >=, the == is good.
			return 0
		}
		if memo[left][right] > 0 {
			return memo[left][right]
		}

		// Need to calculate. Try all of the possible partitions into two:
		for cut := left; cut < right; cut++ {
			// `cut` means dividing [left, right] AFTER the element indexed `cut`.
			// note: there's less-than <, not <=, as even though a <= cut is possible,
			// it is trivially suboptimal (leaves with a score of 0). So skip checking that one.

			sumLeft := prefixSum[cut+1] - prefixSum[left]     // == sum(stoneValue[left .. cut])
			sumRight := prefixSum[right+1] - prefixSum[cut+1] // == sum(stoneValue[cut+1 .. right])

			// Calculate hypothetical scenario A: Bob takes away the right hand side partition
			scorePickLeftPart := sumLeft + dp(left, cut)

			// Calculate hypothetical scenario B: Bob takes away the left hand side partition
			scorePickRightPart := sumRight + dp(cut+1, right)

			// Now actually resolve whatever happens (depending on the actual sums of portitions):
			if sumLeft < sumRight {
				// Bob forces the right part to be taken away, Alice adds the left part's sum
				// to her score, the game continues with [left, cut].
				memo[left][right] = max(memo[left][right], scorePickLeftPart)
			} else if sumLeft > sumRight {
				// Bob forces the left part to be taken away, Alice adds the right part's sum
				// to her score, the game continues with [cut+1, right].
				memo[left][right] = max(memo[left][right], scorePickRightPart)
			} else {
				// The interesting part, and the reason both results need to be calculated before this if.
				// Alice picks whichever part (lhs or rhs) yields the better result. The game continues.
				memo[left][right] = max(memo[left][right], max(scorePickLeftPart, scorePickRightPart))
			}
		}
		return memo[left][right]
	}

	return dp(0, L-1)
}

/*
	dump_memo := func() {
		for row := range L {
			fmt.Println(memo[row])
		}
	}
*/
