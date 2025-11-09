package lc2169

// Runtime complexity: O(log(n)), where n is min(num1, num2)
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2025-11-09
func countOperations(num1 int, num2 int) int {
	ans := 0

	for num1 > 0 && num2 > 0 {
		if num1 < num2 {
			num2, num1 = num1, num2
		}
		ans += num1 / num2
		num1 %= num2
	}

	return ans
}
