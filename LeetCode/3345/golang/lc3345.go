package lc3345

// Runtime complexity: O(n), but n is at most some tens of numbers, so O(1) is arguably correct too.
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-08-06
func smallestNumber(n int, t int) int {
	// Return the smallest number greater than or equal to `n` such that
	// the product of its digits is divisible by t.
	for true {
		dp := digitProduct(n)
		if dp%t == 0 {
			return n
		}
		n += 1
	}
	return -42 // just to fix the "missing return" compiler error.
}

func digitProduct(n int) int {
	ans := 1
	for n > 0 {
		//digit := n % 10
		ans *= (n % 10)
		n /= 10
	}
	return ans
}
