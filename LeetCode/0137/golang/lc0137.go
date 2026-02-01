package lc0137

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: medium+
// Solved on: 2026-02-01
func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones ^= (num & ^twos)
		twos ^= (num & ^ones)
	}
	return ones

	/* Works in C++, does not in Go. Would have to flip between ints and int64s.
	ans := 0
	for i := range 32 {
		accum := 0
		for _, num := range nums {
			// extract ith bit and add it to the sum:
			accum += num >> i & 1
		}
		// sum bits, but modulo three, so that all the numbers that are
		// present a multiple of three times (three included) are not counted.
		accum %= 3
		ans |= accum << i
	}
	*/
}
