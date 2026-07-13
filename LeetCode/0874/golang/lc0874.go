package lc0874

// Runtime complexity: O(len(commands) + len(obstacles))
// Auxiliary space: O(len(obstacles))
// Subjective level: medium+
// Solved on: 2026-07-13
func robotSim(commands []int, obstacles [][]int) int {
	DIRS := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	// isObstacle[c] is an abstraction of a set for storing/querying
	// whether a complex point of [x + iy] is an obstacle.
	isObstacle := map[complex128]bool{}
	for _, o := range obstacles {
		complexPos := complex(float64(o[0]), float64(o[1]))
		isObstacle[complexPos] = true
	}
	ans := 0
	dir := 1 // i.e. "north" {dx=0, dy=1}
	x, y := 0, 0

	for _, cmd := range commands {
		switch cmd {
		case -2: // turn-left command
			dir = (dir + 3) % 4 // same as `dir = (dir-1)%4`, but underflow-safe.
		case -1: // turn-right command
			dir = (dir + 1) % 4
		default:
			for range cmd {
				complexDestination := complex(float64(x+DIRS[dir][0]), float64(y+DIRS[dir][1]))
				if isObstacle[complexDestination] {
					break // the `for range cmd`-loop for this movement.
				}
				x += DIRS[dir][0]
				y += DIRS[dir][1]
			}
			ans = max(ans, (x*x + y*y))
		}
	}

	return ans
}
