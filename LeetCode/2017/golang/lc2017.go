package lc2017

import "math"

func gridGame(grid [][]int) int64 {
	W := len(grid[0])
	var ans int64 = math.MaxInt64
	sum0, sum1 := int64(0), int64(0)
	for _, val := range grid[0] {
		sum0 += int64(val)
	}

	for goDownAt := 0; goDownAt < W; goDownAt++ {
		sum0 -= int64(grid[0][goDownAt])
		ans = min(max(sum0, sum1), ans)
		sum1 += int64(grid[1][goDownAt])
	}
	return ans
}
