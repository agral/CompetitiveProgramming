package lc3197

import "math"

// Runtime complexity: O(H*W)
// Auxiliary space: O(1)
// Subjective level: hard.
func minimumSum(grid [][]int) int {
	H := len(grid)
	W := len(grid[0])
	ans := H * W

	getMinArea := func(h_start int, h_end int, w_start int, w_end int) int {
		y_low := math.MaxInt32
		y_high := 0
		x_low := math.MaxInt32
		x_high := 0
		for y := h_start; y <= h_end; y++ {
			for x := w_start; x <= w_end; x++ {
				if grid[y][x] == 1 {
					y_low = min(y_low, y)
					y_high = max(y_high, y)
					x_low = min(x_low, x)
					x_high = max(x_high, x)
				}
			}
		}
		// Note: getMinArea() might be called over an area that does not contain any "1".
		// Need to handle this case too; then the area is exactly 0.
		if y_low == math.MaxInt32 {
			return 0
		}
		return (1 + y_high - y_low) * (1 + x_high - x_low)
	}

	// -LEFT+MID+RIGHT-
	for w1 := range W {
		for w2 := w1 + 1; w2 < W; w2++ {
			left := getMinArea(0, H-1, 0, w1)
			mid := getMinArea(0, H-1, w1+1, w2)
			right := getMinArea(0, H-1, w2+1, W-1)
			ans = min(ans, left+mid+right)
		}
	}

	// -TOP-
	// -MID-
	// -BOT-
	for h1 := range H {
		for h2 := h1 + 1; h2 < H; h2++ {
			top := getMinArea(0, h1, 0, W-1)
			mid := getMinArea(h1+1, h2, 0, W-1)
			bot := getMinArea(h2+1, H-1, 0, W-1)
			ans = min(ans, top+mid+bot)
		}
	}

	// ----TOP-----
	// -LEFT+RIGHT-
	for h := range H {
		top := getMinArea(0, h, 0, W-1)
		for w := range W {
			left := getMinArea(h+1, H-1, 0, w)
			right := getMinArea(h+1, H-1, w+1, W-1)
			ans = min(ans, top+left+right)
		}
	}

	// -LEFT+RIGHT-
	// ----BOT-----
	for h := range H {
		bot := getMinArea(h, H-1, 0, W-1)
		for w := range W {
			left := getMinArea(0, h-1, 0, w)
			right := getMinArea(0, h-1, w+1, W-1)
			ans = min(ans, bot+left+right)
		}
	}

	// -LEFT+TOPRIGHT-
	// -LEFT+BOTRIGHT-
	for w := range W {
		left := getMinArea(0, H-1, 0, w)
		for h := range H {
			topright := getMinArea(0, h, w+1, W-1)
			botright := getMinArea(h+1, H-1, w+1, W-1)
			ans = min(ans, left+topright+botright)
		}
	}

	// -TOPLEFT+RIGHT-
	// -BOTLEFT+RIGHT-
	for w := range W {
		right := getMinArea(0, H-1, w, W-1)
		for h := range H {
			topleft := getMinArea(0, h, 0, w-1)
			botleft := getMinArea(h+1, H-1, 0, w-1)
			ans = min(ans, right+topleft+botleft)
		}
	}

	return ans
}
