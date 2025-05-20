package lc3355

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// where n is the length of nums slice.
func isZeroArray(nums []int, queries [][]int) bool {
	line := make([]int, len(nums)+1)
	decr := 0

	for _, query := range queries {
		l := query[0]
		r := query[1]
		line[l] += 1
		line[r+1] -= 1
	}

	for i := range nums {
		decr += line[i]
		if decr < nums[i] {
			return false
		}
	}
	return true
}
