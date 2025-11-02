package lc2257

// Runtime complexity: O(W*H)
// Auxiliary space: O(W*H)
// Subjective level: medium

// Note: using bytes instead of ints actually saves memory.
const UNGUARDED = byte(0)
const GUARDED = byte(1)
const GUARD = byte(2)
const WALL = byte(3)

func MakeByteGrid(H int, W int) [][]byte {
	grid := make([][]byte, H)
	for h := range H {
		grid[h] = make([]byte, W)
	}
	return grid
}

// Note: can't do it with a standard 2D grid approach; the constraints are too lax.
func countUnguarded(H int, W int, guards [][]int, walls [][]int) int {
	grid := MakeByteGrid(H, W)
	up := MakeByteGrid(H, W)
	right := MakeByteGrid(H, W)
	down := MakeByteGrid(H, W)
	left := MakeByteGrid(H, W)

	for _, wall := range walls {
		grid[wall[0]][wall[1]] = WALL
	}
	for _, guard := range guards {
		grid[guard[0]][guard[1]] = GUARD
	}

	// fill up the *left* grid
	for h := range H {
		last := byte(0) // whatever was seen last; at first nothing; but then remembers either a GUARD or a WALL.
		// won't ever return to UNGUARDED once a GUARD or a WALL has been seen.
		// will update whenever a GUARD or a WALL is seen.
		for w := range W {
			if grid[h][w] >= GUARD { // either a guard or a wall
				last = grid[h][w]
			} else {
				left[h][w] = last
			}
		}
	}

	// fill up the *right* grid
	for h := range H {
		last := byte(0)
		for w := W - 1; w >= 0; w-- {
			if grid[h][w] >= GUARD { // either a guard or a wall
				last = grid[h][w]
			} else {
				right[h][w] = last
			}
		}
	}

	// fill up the *up* grid
	for w := range W {
		last := byte(0)
		for h := range H {
			if grid[h][w] >= GUARD { // either a guard or a wall
				last = grid[h][w]
			} else {
				up[h][w] = last
			}
		}
	}

	// fill up the *down* grid
	for w := range W {
		last := byte(0)
		for h := H - 1; h >= 0; h-- {
			if grid[h][w] >= GUARD { // either a guard or a wall
				last = grid[h][w]
			} else {
				down[h][w] = last
			}
		}
	}

	// With these 4-directional extra info it's immediately known for each field
	// whether a guard is overseeing it from any of the four directions:
	ans := 0
	for h := range H {
		for w := range W {
			if (grid[h][w] == 0) && (up[h][w] != GUARD) && (right[h][w] != GUARD) &&
				(down[h][w] != GUARD) && (left[h][w] != GUARD) {
				ans += 1
			}
		}
	}
	return ans
}
