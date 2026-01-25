package lc0060

import (
	"strconv"
	"strings"
)

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: hard
// Solved on: 2026-01-25
func getPermutation(n int, k int) string {
	// 1. Build the answer string efficiently:
	ans := strings.Builder{}

	// 2. Make k zero-indexed, it is 1-indexed in the problem statement:
	k -= 1

	// 3. Prepare the list of digits to be used in the answer sequence:
	// Note: 1 <= n <= 9, so it's guaranteed that digits are digits not exceeding 9.
	digits := make([]int, n)
	for i := range n {
		digits[i] = i + 1
	}

	// 4. Prepare the list of factorials of numbers from 0 to n, inclusively:
	factorials := make([]int, n+1)
	factorials[0] = 1
	factorials[1] = 1
	for i := 2; i <= n; i++ {
		factorials[i] = factorials[i-1] * i
	}

	// 5. Generate the kth permutation of the sequence:
	// Key insight: the set of n digits has n! possible permutations.
	// So for each n-long permutation list, the first digit stays the same,
	// and subsequent (n-1)-long permutation "kind-of-repeats" n times,
	// each time with different, but clearly defined, digits.
	// For example, with n=4, it starts with (on the right, n=3):
	// Note that the numbers on the right equal three last numbers on the left,
	// plus one each:
	// 1 2 3 4             1 2 3
	// 1 2 4 3             1 3 2
	// 1 3 2 4             2 1 3
	// 1 3 4 2             2 3 1
	// 1 4 2 3             3 1 2
	// 1 4 3 2             3 2 1
	// 2 1 3 4             (end-of-sequence for n=3)
	// Key insight #2: the first number of n-long sequence stays the same (n-1)! times,
	// then switches to a +1, which will also stay the same (n-1)! times.
	// All the numbers in [1..n] have to be used, each exactly once.
	// So, nth number can be generated in O(1) using precomputed factorials.
	for i := n - 1; i >= 0; i-- {
		d := k / factorials[i]
		k %= factorials[i]

		// maybe could have been done faster using byte values; don't care. Like this approach.
		ans.WriteString(strconv.Itoa(digits[d]))

		// Once the digit has been used, remove it from the list of usable digits:
		digits = append(digits[:d], digits[d+1:]...)
	}

	return ans.String()
}
