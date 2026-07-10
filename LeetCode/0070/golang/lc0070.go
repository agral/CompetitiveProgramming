package lc0070

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2026-07-10
func climbStairs(n int) int {
	// base cases: staircase of 1 total steps: 1 way to reach the top
	//             staircase of 2 total steps: 2 ways to reach the top.
	// and the total number of ways to reach n stairs is always ways(n-1) + ways(n-2).
	// That's a Fibonacci's sequence!
	fib := 1     // holds Fibonacci's kth number, we're calculating the nth one.
	fibPrev := 1 // holds fib_{k-1}
	fibAnte := 0 // holds fib_{k-2}
	for range n {
		fib = fibPrev + fibAnte
		fibAnte = fibPrev
		fibPrev = fib
	}

	return fib
}
