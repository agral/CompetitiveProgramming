package lc3483

// Runtime complexity: O(10*10*10) == O(1)
// Auxiliary space: O(10*10*10) == O(1)
// Subjective level: easy+
// Solved on: 2026-09-11
func totalNumbers(digits []int) int {
	L := len(digits) // L is guaranteed to be between 3 and 10, inclusively.
	valid_nums := map[int]bool{}

	for f := range L {
		// The first digit can be anything but zero.
		if digits[f] != 0 {
			for s := range L {
				// The second digit can be anything, but not the first one.
				if s != f {
					for t := range L {
						// The third digit must be even (as the resulting number has to be even).
						// and not either the first or second digit - these are already taken.
						if t != f && t != s && digits[t]%2 == 0 {
							// This number is valid - add it to the set of valid numbers.
							// That number might be repeated, so don't blindly increase a counter,
							// instead hold all the valid numbers in a set (represented in Go as map[int]bool).
							num := (100 * digits[f]) + (10 * digits[s]) + digits[t]
							valid_nums[num] = true
						}
					}
				}
			}
		}
	}
	return len(valid_nums)
}
