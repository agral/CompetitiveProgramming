package lc2946

// Runtime complexity: O(H*W)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-03-27
func areSimilar(mat [][]int, k int) bool {
	H := len(mat)
	W := len(mat[0])
	K := k % W
	for h := range H {
		direction := -1 // even indexed rows are shifted cyclically to the left (dir=-1)
		if h%2 == 1 {
			direction = 1 // odd indexed rows are shifted cyclically to the right (dir=1)
		}
		for w := range W {
			shifted := ((w + W) + (K * direction)) % W
			if mat[h][w] != mat[h][shifted] {
				return false
			}
		}
	}
	return true
}
