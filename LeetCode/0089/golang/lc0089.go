package lc0089

import "math"

func powTwo(n int) int {
	return int(math.Pow(2.0, float64(n)))
}

// Runtime complexity: O(2^n)
// Auxiliary space: O(2^n)
// Subjective level: medium
// Solved on: 2026-01-31
func grayCode(n int) []int {
	// Note: I'm actually generating a binary-reflected Gray code of length 2^n.
	// (it is permitted and probably expected by the problem statement).
	// Idea and algorithm description:
	// https://en.wikipedia.org/wiki/Gray_code#Constructing_an_n-bit_Gray_code
	// (but I'm optimizing this by preallocating the entire slice
	// and keeping track of the total count of already generated entries -
	// with this trick I don't have to allocate extra memory for intermediate slices).
	L := powTwo(n)
	ans := make([]int, L)
	ans[0] = 0
	numGen := 1
	for i := range n {
		for j := numGen - 1; j >= 0; j -= 1 {
			ans[numGen] = ans[j] | 1<<i
			numGen += 1
		}
	}
	return ans
}
