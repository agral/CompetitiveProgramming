package lc2094

import "unsafe"

// Converts `false` to 0, `true` to 1. Golang has no automatic conversion of such kind.
// Also, is pretty fast.
func fastBoolToInt(b bool) int {
	return int(*(*byte)(unsafe.Pointer(&b)))
}

// Runtime complexity: O(10^3) == O(1)
// Auxiliary space: O(|digits| + 10)
func findEvenNumbers(digits []int) []int {
	ans := []int{}
	count := make([]int, 10)
	for _, d := range digits {
		count[d]++
	}

	// If _xyz_ can be constructed, add it to the answer list.
	// The generated numbers are even and in ascending order.
	for x := 1; x <= 9; x++ {
		for y := 0; y <= 9; y++ {
			for z := 0; z <= 8; z += 2 {
				if count[x] > 0 &&
					count[y] > fastBoolToInt(x == y) &&
					count[z] > fastBoolToInt(x == z)+fastBoolToInt(y == z) {
					ans = append(ans, 100*x+10*y+z)
				}
			}
		}
	}
	return ans
}
