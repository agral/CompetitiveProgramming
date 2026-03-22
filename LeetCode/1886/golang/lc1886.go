package lc1886

import "slices"

// Runtime complexity: O(N^2)
// Auxiliary space: O(1) as I'm reusing the provided `mat`; O(N^2) in the general case.
// Subjective level: medium-. Easy+ if you know the trick to rotate a 2D slice.
// Solved on: 2026-03-22
func findRotation(mat [][]int, target [][]int) bool {
	N := len(mat)

	// checks if 2d slice `mat` is equal to `target`.
	// skips sanity checks; it is guaranteed that both matrices are of size N by N.
	isMatching := func() bool {
		for y := range N {
			if !slices.Equal(mat[y], target[y]) {
				return false
			}
		}
		return true
	}

	// 0 degrees rotation:
	if isMatching() {
		return true
	}

	// rotates the `mat` 2d slice 90 degrees clockwise.
	rotate := func() {
		for y := range N / 2 {
			mat[y], mat[N-1-y] = mat[N-1-y], mat[y]
		}
		for y := 0; y < N; y++ {
			for x := y + 1; x < N; x++ {
				mat[y][x], mat[x][y] = mat[x][y], mat[y][x]
			}
		}
	}

	for range 3 { // 90, 180, 270 degrees rotations:
		rotate()
		if isMatching() {
			return true
		}
	}

	// no match after initial check + 3 rotations; so it can't be done:
	return false
}
