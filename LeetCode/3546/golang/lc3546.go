package lc3546

// Runtime complexity: O(H*W)
// Auxiliary space: O(H+W)
// Subjective level: medium-
// Solved on: 2026-03-25
func canPartitionGrid(grid [][]int) bool {
	H := len(grid)
	W := len(grid[0])

	// 1. Calculate sum of row-wise and column-wise sums,
	// and the total sum of the entire grid:
	rowSums := make([]int, H)
	colSums := make([]int, W)
	totalSum := 0
	for h := range H {
		for w := range W {
			rowSums[h] += grid[h][w]
			colSums[w] += grid[h][w]
		}
		totalSum += rowSums[h]
	}

	// 2. Try horizontal (row-wise) partitioning; "cut after h'th row":
	topSum := 0
	bottomSum := totalSum
	for h := 0; h < H-1 && topSum < bottomSum; h++ {
		topSum += rowSums[h]
		bottomSum -= rowSums[h]
		if topSum == bottomSum {
			// found a valid horizontal partitioning!
			return true
		}
	}

	// 3. Try a vertical (column-wise) partitioning; "cut after w'th column":
	leftSum := 0
	rightSum := totalSum
	for w := 0; w < W-1 && leftSum < rightSum; w++ {
		leftSum += colSums[w]
		rightSum -= colSums[w]
		if leftSum == rightSum {
			// found a valid vertical partitioning!
			return true
		}
	}
	return false
}
