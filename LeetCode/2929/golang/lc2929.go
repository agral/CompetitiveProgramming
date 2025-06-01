package lc2929

// Runtime complexity: O(1)
// (that's right: n_choose_k has a for loop, but k is 2, so there's just two iterations).
// Auxiliary space: O(1)
func distributeCandies(n int, limit int) int64 {
	limit_plus_one := limit + 1

	waysNoLimit := getNumWays3(n)
	threeChildrenExceedLimit := getNumWays3(n - 3*limit_plus_one)
	twoChildrenExceedLimit := getNumWays3(n - 2*limit_plus_one)
	oneChildExceedsLimit := getNumWays3(n - limit_plus_one)

	return waysNoLimit - threeChildrenExceedLimit - 3*(oneChildExceedsLimit-twoChildrenExceedLimit)
}

// Returns the value of n-choose-k, calculated using the multiplicative formula:
// https://en.wikipedia.org/wiki/Binomial_coefficient
func n_choose_k(n int, k int) int64 {
	ans := int64(1)
	for i := 1; i <= k; i++ {
		ans = ans * int64(n-i+1) / int64(i)
	}
	return ans
}

// Returns the number of ways to distribute `n` candies to exactly 3 children,
// not taking any limit into account.
func getNumWays3(n int) int64 {
	if n < 0 {
		return 0
	}
	return n_choose_k(n+2, 2)
}
