package lc3742

// Runtime complexity: O(H*W*k)
// Auxiliary space: O(H*W*h)
// Subjective level: hard
// Solved on: 2026-04-30
func maxPathScore(grid [][]int, k int) int {
	// 1. Prepare the score array (DP).
	//    dp[h][w][cost] := maximum score that can be obtained at grid[h][w],
	//    using exactly `cost` total cost.
	//    note: now switched to using `cost` as remaining cost to be used.
	const INVALID = -1
	const NEGATIVE_INFINITY = -(1 << 30)
	H := len(grid)
	W := len(grid[0])
	dp := make([][][]int, H)
	for h := range H {
		dp[h] = make([][]int, W)
		for w := range W {
			dp[h][w] = make([]int, k+1)
			for i := range k + 1 {
				dp[h][w][i] = INVALID
			}
		}
	}

	var getScore func(h int, w int, remainingCost int) int
	getScore = func(h int, w int, remainingCost int) int {
		if h < 0 || w < 0 || remainingCost < 0 {
			// don't simply return a -1, as this will be drowned by the positive results.
			// return a big negative value, so that whatever was added to it,
			// the end result is still negative. Handle specially in the end.
			return NEGATIVE_INFINITY
		}
		// at this point it is guaranteed that remainingCost >= 0, so the path is valid/reachable.
		if h == 0 && w == 0 {
			return 0 // anyway, it is guaranteed by the constraints too.
		}

		// if already calculated, return the memoized value:
		if dp[h][w][remainingCost] != INVALID {
			return dp[h][w][remainingCost]
		}

		newRemainingCost := remainingCost
		if grid[h][w] > 0 {
			newRemainingCost -= 1
		}
		ans := grid[h][w]
		checkFromTop := getScore(h-1, w, newRemainingCost)
		checkFromLeft := getScore(h, w-1, newRemainingCost)
		ans += max(checkFromTop, checkFromLeft)

		// memoize the cost to this cell (at this point guaranteed optimal!), then return:
		dp[h][w][remainingCost] = ans
		return ans
	}

	ans := getScore(H-1, W-1, k)
	if ans < 0 {
		return INVALID
	}
	return ans
}
