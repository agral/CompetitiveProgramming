package lc0840

import "slices"

func permute(seq []int) []int {
	// permute the sequence forward.
	// This will not be called in the final answer, I don't care about performance optimization.

	// 1. find the last pair of digits forming an increasing sequence;
	//    e.g. '34' in 1234, '24' in '1243', '24' in '1324', '34' in '1342', and so on.
	L := len(seq)
	i := L - 2

	// dirty hack: this won't ever be called on a reverse-sorted sequence,
	// so it's guaranteed that seq[i] < seq[i+1] for at least one i >= 0.
	for seq[i] > seq[i+1] {
		i--
	}
	// now i contains the index of the first digit in the last increasing pair of digits of seq.

	// 2. find the lowest digit in seq[i+1] .. seq[7], that is greater than seq[i].
	iEx := i + 1
	for k := i + 2; k < L; k++ {
		if seq[i] < seq[k] && seq[k] < seq[iEx] {
			iEx = k
		}
	}

	// 3. swap seq[i] and seq[iEx].
	seq[i], seq[iEx] = seq[iEx], seq[i]

	// 4. Sort seq[i+1] .. seq[L-1] in ascending order.
	for o := i + 1; o < L-1; o++ {
		for u := o + 1; u < L; u++ {
			if seq[o] > seq[u] {
				seq[o], seq[u] = seq[u], seq[o]
			}
		}
	}
	return seq
}

// Unused in the final solution, the result of calling this is hardcoded.
func generate3x3MagicSquares() [][]int {
	// Magic square of 3x3 has to have '5' in its center.
	// Omit the central 5 in a "work" slice. It's always in the same place anyway.
	// So the magic square maps to the following slice,
	// with the 'F' in the centre ignored, that's always equal five (5):
	// 0 1 2
	// 3 F 4 ==> []int{0, 1, 2, 3, 4, 5, 6, 7}.
	// 5 6 7
	work := []int{1, 2, 3, 4, 6, 7, 8, 9}
	ans := [][]int{}
	for !slices.Equal(work, []int{9, 8, 7, 6, 4, 3, 2, 1}) {
		// check if this is a magic square. If so, add it to the answer list.
		if (work[0]+work[1]+work[2] == 15) &&
			(work[3]+work[4] == 10) &&
			(work[5]+work[6]+work[7] == 15) &&
			(work[0]+work[3]+work[5] == 15) &&
			(work[1]+work[6] == 10) &&
			(work[2]+work[4]+work[7] == 15) &&
			(work[0]+work[7] == 10) &&
			(work[2]+work[5] == 10) {
			safecopy := make([]int, len(work))
			copy(safecopy, work)
			ans = append(ans, safecopy)
		}

		// permute the square forward:
		work = permute(work)
	}

	return ans
}

// Runtime complexity: O(n^2). Optimal, as the grid dimensions are not known at compile time,
// but the length of MAGIC (8) is known.
// Auxiliary space: O(8*8) == O(1). Surely.
// Subjective level: medium+ or even hard, mostly due to a lot of corner cases to consider;
// while the final solution indeed looks like "magic".
// A variation of such challenge may appear on the SW Test.
// Solved on: 2025-12-30
func numMagicSquaresInside(grid [][]int) int {
	// These are all the possible positions of numbers in a magic square,
	// the 5 not included is always present in the centre.
	// Hardcoded instead of calling `generate3x3MagicSquares()` to shave off some milliseconds.
	MAGIC := [][]int{
		{2, 7, 6, 9, 1, 4, 3, 8},
		{2, 9, 4, 7, 3, 6, 1, 8},
		{4, 3, 8, 9, 1, 2, 7, 6},
		{4, 9, 2, 3, 7, 8, 1, 6},
		{6, 1, 8, 7, 3, 2, 9, 4},
		{6, 7, 2, 1, 9, 8, 3, 4},
		{8, 1, 6, 3, 7, 4, 9, 2},
		{8, 3, 4, 1, 9, 6, 7, 2},
	} // also fun fact, all these squares have even values at even positions, and odd values at odd positions!
	Y := len(grid)
	X := len(grid[0])
	ans := 0

	for y := 0; y < Y-2; y++ {
		for x := 0; x < X-2; x++ {
			if grid[y+1][x+1] == 5 {
				// check if this square is magic:
				for _, m := range MAGIC {
					if grid[y][x] == m[0] && grid[y][x+1] == m[1] && grid[y][x+2] == m[2] &&
						grid[y+1][x] == m[3] && grid[y+1][x+2] == m[4] &&
						grid[y+2][x] == m[5] && grid[y+2][x+1] == m[6] && grid[y+2][x+2] == m[7] {
						ans += 1
						break
					}
				}
			}
		}
	}

	return ans
}
