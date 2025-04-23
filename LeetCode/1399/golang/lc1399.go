package lc1399

// Runtime complexity: O(n logn)
//   - getDecimalDigitsSum has O(log10n)
//   - and it is executed for every k in range 1 to n.
//   - arguing that it's getDecimalDigitsSum is constant since it's at most 5 digits
//     in the given range is not wrong (but it would be in the general case of unbounded n).
//
// Auxiliary space: O(37) == O(1).
// - In the general case of unbounded n, it's O(9*log10(n)) == O(log10(n))
func countLargestGroup(n int) int {
	// 1 <= n <= 10^4.
	// The minimum and maximum possible digit sum for numbers in this range is 1 and 36 (==9+9+9+9), respectively.
	counts := make([]int, 36+1) // +1 to account for zero-indexing.

	for k := 1; k <= n; k++ {
		counts[getDecimalDigitsSum(k)]++
	}

	largest_group_count := 0
	largest_group_length := 0
	for _, v := range counts {
		if v > largest_group_length {
			largest_group_length = v
			largest_group_count = 1
		} else if v == largest_group_length {
			largest_group_count++
		}
	}
	return largest_group_count
}

func getDecimalDigitsSum(n int) int {
	ans := 0
	for n > 0 {
		ans += n % 10
		n /= 10
	}
	return ans
}
