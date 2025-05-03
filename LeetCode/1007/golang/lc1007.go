package lc1007

// Runtime complexity: O(n)
// Auxiliary space: O(3*7) == O(1)
func minDominoRotations(tops []int, bottoms []int) int {
	SZ := len(tops)
	count_of_tops := make([]int, 7)
	count_of_bottoms := make([]int, 7)
	count_of_same := make([]int, 7)

	for i := range SZ {
		count_of_tops[tops[i]] += 1
		count_of_bottoms[bottoms[i]] += 1
		if tops[i] == bottoms[i] {
			count_of_same[tops[i]] += 1
		}
	}

	for i := 1; i <= 6; i++ {
		if count_of_tops[i]+count_of_bottoms[i]-count_of_same[i] == SZ {
			return SZ - max(count_of_tops[i], count_of_bottoms[i])
		}
	}

	return -1
}
