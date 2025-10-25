package lc1716

/**
 * approach: calculate sums accumulated during whole weeks,
 * then add sum accumulated in the last non-full week (0 to 6 days).
 * 1st week:  1 + 2 + 3 + 4 + 5 + 6 + 7 == 4 * 7 == 28
 * 2nd week:  2 + 3 + 4 + 5 + 6 + 7 + 8 == 5 * 7 == 35
 * 3rd week:  (...)                     == 6 * 7 == 42
 * nth week:  (...)                     == (n + 3) * 7.
 * Sum accumulated during full weeks: (4 + 5 + 6 + ... + (n+3)) * 7.
 *
 * Then, in the last non-full week (0 to 6 days):
 * k + 1 + k + 2 + k + 3 + ... + k + 6,
 * where k is num_weeks, and the last term is k + n % 7.
 */

// Runtime complexity: O(1)
// Auxiliary space: O(1)
// Subjective level: easy.
func totalMoney(n int) int {
	COUNT_FULL_WEEKS := n / 7
	COUNT_EXTRA_DAYS := n % 7
	accumulated_full_weeks := 7 * arithmeticSum(4, COUNT_FULL_WEEKS+3)
	accumulated_last_week := arithmeticSum(COUNT_FULL_WEEKS+1, COUNT_FULL_WEEKS+COUNT_EXTRA_DAYS)
	return accumulated_full_weeks + accumulated_last_week
}

// Calculates the sum of all ints between low and high, i.e. <low, low+1, ..., high-1, high>
func arithmeticSum(low int, high int) int {
	return (high + low) * (high - low + 1) / 2
}
