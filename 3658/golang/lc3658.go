package lc3658

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Runtime complexity: O(1)
// Auxiliary space: O(1)
// Subjective level: easy & satisfying
// Solved on: 2026-07-15
func gcdOfOddEvenSums(n int) int {
	// Calculate the sums of the first n even and the first n odd numbers.
	// Note: for the first n even numbers, let's take n=5 as in Example 2 of the task's description.
	// sumEven5 = sum{2, 4, 6, 8, 10}. But note that each number here is even, so 2 can be factored
	// outside of the sum: sumEven5 = 2 * sum{1, 2, 3, 4, 5}. And sum{1, 2, 3, 4, 5} is the sum
	// of first n (5) natural numbers, calculable by the easy formula n(n+1)/2, here 5(5+1)/2 == 15.
	// So sumEvenN is twice that, i.e. sumEvenN == n(n+1).
	// And for the sumOddN, it's easy to see that each number there corresponds to another one from
	// the sumEvenN - but is smaller by 1 from its counterpart in sumEvenN. For n=5,
	// sumOddN == {1, 3, 5, 7, 9}
	// sumEvenN  == {2, 4, 6, 8, 10}
	// Because of that, it's easy to induce that sumOddN = sumEvenN - n, as there are n elements,
	// each smaller by one than the ones that sum to sumEvenN.

	// So, in the end:
	// sumEvenN := n * (n+1)
	// sumOddN := (n * (n+1)) - n == n * n == n^2
	// Alternatively, and I'm going to continue with it: sumOdd := n*n, sumEven = sumOdd + n.

	// given that we're looking for a GCD of (n*n, n*(n+1)), it's going to be n, in every case.
	// That gcd function I've expanded from my snippets above - not even needed.
	// It's indeed an `easy` one, once you work it out.

	return n
}
