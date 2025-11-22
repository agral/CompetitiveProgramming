package lc3190

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: (very) easy
// Solved on: 2025-11-22
func minimumOperations(nums []int) int {
	// This question boils down to counting the numbers that are not divisible by 3.
	// This is since every number has to be either 0, 1, or 2 in modulo 3.
	// Numbers that are ==1 (mod 3) need to have one subtract-operation performed to be 0 (mod 3);
	// while numbers that are ==2 (mod 3) are made 0 (mod 3) after one add-operation.
	// Of course, numbers that are already ==0 (mod 3) require no operations.
	ans := 0
	for _, num := range nums {
		if num%3 > 0 {
			ans += 1
		}
	}
	return ans
}
