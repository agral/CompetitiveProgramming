package lc2211

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2025-12-04
func countCollisions(directions string) int {
	isLeftBlocked := false
	isRightBlocked := false
	numCollisions := 0

	// 1st pass: go left-to-right, count the number of left-facing cars that collide.
	for i := 0; i < len(directions); i++ {
		if directions[i] == 'L' {
			if isLeftBlocked {
				numCollisions += 1
			} // else this car escapes to the left, do nothing.
		} else { // encountered either a right-facing car, or a stopped car. The following L-cars can not escape.
			isLeftBlocked = true
		}
	}

	// 2nd pass: go right-to-left, count the number of right-facing cars that collide.
	for i := len(directions) - 1; i >= 0; i-- {
		if directions[i] == 'R' {
			if isRightBlocked {
				numCollisions += 1
			} // else this car escapes to the right, do nothing.
		} else { // either left-facing or a stopped car. The previous R-facing cars will collide.
			isRightBlocked = true
		}
	}
	return numCollisions
}
