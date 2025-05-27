package lc2894

// Runtime complexity: O(1)
// Auxiliary space: O(1)

// This is a math question, not a programming question.
// But a fun change from the usual thing.

func sumOfDivisibleIntegers(n int, m int) int {
	var last_divisor int = (n / m) * m
	if last_divisor == 0 {
		return 0
	}
	first_divisor := m
	num_divisors := (last_divisor-first_divisor)/m + 1
	return num_divisors * (first_divisor + last_divisor) / 2
}

func differenceOfSums(n int, m int) int {
	var sum_n int = n * (n + 1) / 2
	sumDivisors := sumOfDivisibleIntegers(n, m)
	sumNonDivisors := sum_n - sumDivisors
	return sumNonDivisors - sumDivisors
}
