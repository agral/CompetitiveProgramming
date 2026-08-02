package lc0877

// Runtime complexity: O(1)
// Auxiliary space: O(1)
// Subjective level: medium.
// Solved on: 2026-08-02
func stoneGame(piles []int) bool {
	// Given that the number of piles is even (so important!),
	// Alice can always win, given that she plays optimally.
	// This is because she can in effect: select the *best* pile, on move 1.
	// Bob would have a chance at winning only at "impossible to win" scenarios,
	// such as []int{1, 101, 1}; but these are not allowed as the total piles' count is even!
	// So in effect, Alice wins in every scenario.
	return true
}
