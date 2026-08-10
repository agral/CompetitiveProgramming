package lc1510

// Runtime complexity: O(n * sqrt(n))
// Auxiliary space: O(n)
// Subjective level: hard.
// Solved on: 2026-08-10
func winnerSquareGame(n int) bool {
	isWinner := make([]bool, n+1)
	for k := 1; k <= n; k++ {
		for sq := 1; sq*sq <= k; sq++ {
			if !isWinner[k-(sq*sq)] {
				// just found out that removing sq*sq stones makes Alice win with perfect play.
				isWinner[k] = true
				break // the outer, for-sq loop. A winning strategy is found, so no need to iterate further.
			}
		}
	}

	return isWinner[n]
}
