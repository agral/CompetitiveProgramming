package lc0657

// Runtime complexity: O(n)
// Auxiliary space: O(4) == O(1)
// Subjective level: easy
// Solved on: 2026-04-05
func judgeCircle(moves string) bool {
	l := 0
	r := 0
	u := 0
	d := 0
	// could also have been made simpler with two vars; "x" and "y"
	// (x: decrement on "l", increment on "r"); similar with u/d for y.
	// but it's an easy leetcode problem, let's not overengineer it.
	for _, ch := range moves {
		switch ch {
		case 'D':
			d += 1
		case 'L':
			l += 1
		case 'R':
			r += 1
		case 'U':
			u += 1
		}
	}
	// when the total count of down-moves equals the count of up-moves,
	// and the count of left-moves equals the tally of right-moves,
	// then the robot ends up whereever it started (i.e. at the origin).
	return d == u && l == r
}
