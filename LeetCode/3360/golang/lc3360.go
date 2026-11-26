package lc3360

// Runtime complexity: O(1)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-11-26
func canAliceWin(n int) bool {
	// 1 <= n <= 50.
	turn := true // true: Alice's turn; false: Bob's turn
	numToRemove := 10
	for n >= numToRemove {
		n -= numToRemove
		numToRemove -= 1
		turn = !turn
	}
	return turn == false // Alice wins if it's Bob's turn when no further move can be made.
}
