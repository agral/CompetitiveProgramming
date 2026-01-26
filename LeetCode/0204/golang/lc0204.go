package lc0204

import "math"

// Runtime complexity: O(2n) == O(n).
//   - first pass over the sieve to mark composite numbers
//   - second pass over the sieve to count the primes left at the end.
//
// Auxiliary space: O(n)
// Subjective level: medium
// Solved on: 2026-01-26
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	// Note: sieve is of size n, since we want the count of prime numbers strictly less than n.
	// that makes the sieve contain indices from 0 to n-1.
	// Want to keep the numbers 0 and 1 as not prime, and 2 to (n-1) as prime:
	sieve := make([]bool, n) // a standard sieve of Eratosthenes.
	for i := 2; i < n; i++ {
		sieve[i] = true
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if sieve[i] { // further optimization: don't re-check numbers that are already known composites.
			for composite := 2 * i; composite < n; composite += i {
				sieve[composite] = false
			}
		}
	}

	// Count the numbers that are still (i.e. actually) prime.
	ans := 0
	for _, val := range sieve {
		if val {
			ans += 1
		}
	}
	return ans
}
