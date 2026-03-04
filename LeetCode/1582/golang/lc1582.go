package lc1582

// Runtime complexity: O(H*W)
// Auxiliary space: O(H+W)
// Subjective level: easy
// Solved on: 2026-03-04
func numSpecial(mat [][]int) int {
	H := len(mat)
	W := len(mat[0])

	sumRows := make([]int, H)
	sumCols := make([]int, W)
	for h := range H {
		for w := range W {
			sumRows[h] += mat[h][w] // can go like this, since mat is a binary matrix
			sumCols[w] += mat[h][w] // (values are either `0` or `1`).
		}
	}

	ans := 0
	for h := range H {
		for w := range W {
			if mat[h][w] == 1 && sumRows[h] == 1 && sumCols[w] == 1 {
				ans += 1
			}
		}
	}

	return ans
}
