package lc2906

// Runtime complexity: O(3*H*W) == O(H*W)
// Auxiliary space: O(2*H*W) == O(H*W)
// Subjective level: medium
// Solved on: 2026-03-24
func constructProductMatrix(grid [][]int) [][]int {
	const MOD int = 12_345
	H := len(grid)
	W := len(grid[0])

	// 1. Prepare the answer matrix:
	ans := make([][]int, H)
	for h := range H {
		ans[h] = make([]int, W)
	}

	// 2. Prepare the prefix and suffix matrices.
	// p[h][w] := product of all the `grid` entries (column-by-column, for each applicable row)
	//            up to but not including grid[h][w]; all modulo MOD
	// s[h][w] := product of all the `grid` entries from grid[h][w+1] to the end;
	//            i.e. product of all the grid entries that follow grid[h][w]; all modulo MOD.
	p := make([][]int, H)
	s := make([][]int, H)
	for h := range H {
		p[h] = make([]int, W)
		s[h] = make([]int, W)
	}

	// 3. Fill up prefix and suffix matrices:
	// note: it's necessary to apply the modulo division twice;
	// as entries in `grid` can be as high as 1e9.
	p[0][0] = 1
	for h := range H {
		if h > 0 { // if h == 0, then p[0][0] == 1 by definition.
			p[h][0] = ((grid[h-1][W-1] % MOD) * p[h-1][W-1]) % MOD
		}
		for w := 1; w < W; w++ {
			p[h][w] = ((grid[h][w-1] % MOD) * p[h][w-1]) % MOD
		}
	}
	s[H-1][W-1] = 1
	for h := H - 1; h >= 0; h-- {
		if h < H-1 { // similarly if h == H-1, then s[H-1][W-1] == 1, by definition.
			s[h][W-1] = ((grid[h+1][0] % MOD) * s[h+1][0]) % MOD
		}
		for w := W - 2; w >= 0; w-- {
			s[h][w] = ((grid[h][w+1] % MOD) * s[h][w+1]) % MOD
		}
	}

	// 4. Fill up the answer array. ans[h][w] := p[h][w] * s[h][w] mod MOD.
	for h := range H {
		for w := range W {
			ans[h][w] = (p[h][w] * s[h][w]) % MOD
		}
	}

	return ans
}
