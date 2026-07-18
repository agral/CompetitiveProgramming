package lc1914

// Runtime complexity: O(k*H*W), generally requires rotating each ring (H*W) k times.
// Auxiliary space: O(H*W) (for storing one ring being rotated)
// Subjective level: medium, feels like too much grinding, would not recommend.
// Solved on: 2026-05-09
func rotateGrid(grid [][]int, k int) [][]int {
	H := len(grid)
	W := len(grid[0])

	rotateRing := func(ringIdx int, k int) {
		// 1. Check whether it's necessary to rotate at all.
		// - it is a NOP when k is equal to the ring's length or a multiple of it.
		// - it is k modulo the ring's length for k > ring's length.
		//L := 2 * ((H - 1 - (2 * ringIdx)) + (W - 1 - (2 * ringIdx))) // the length of ringIdx's flattened ring.

		// 1. "unloop" the ring into a flat, one-dimensional slice:
		flat := []int{} // a "flattened" ring.
		for x := ringIdx; x < W-1-ringIdx; x++ {
			flat = append(flat, grid[ringIdx][x])
		}
		for y := ringIdx; y < H-1-ringIdx; y++ {
			flat = append(flat, grid[y][W-1-ringIdx])
		}
		for x := W - 1 - ringIdx; x > ringIdx; x-- {
			flat = append(flat, grid[H-1-ringIdx][x])
		}
		for y := H - 1 - ringIdx; y > ringIdx; y-- {
			flat = append(flat, grid[y][ringIdx])
		}

		// 2. Check if it's necessary to rotate at all; it is not when k is a multiple
		// of the ring's total length.
		L := len(flat)
		k %= L
		if k == 0 {
			return
		}

		// 3. Perform the rotation by rewriting the flattened array back to the original grid.
		for x := ringIdx; x < W-1-ringIdx; x++ {
			grid[ringIdx][x] = flat[k]
			k = (k + 1) % L
		}
		for y := ringIdx; y < H-1-ringIdx; y++ {
			grid[y][W-1-ringIdx] = flat[k]
			k = (k + 1) % L
		}
		for x := W - 1 - ringIdx; x > ringIdx; x-- {
			grid[H-1-ringIdx][x] = flat[k]
			k = (k + 1) % L
		}
		for y := H - 1 - ringIdx; y > ringIdx; y-- {
			grid[y][ringIdx] = flat[k]
			k = (k + 1) % L
		}
	}

	numRings := min(H/2, W/2)
	for i := range numRings {
		rotateRing(i, k)
	}
	return grid
}
