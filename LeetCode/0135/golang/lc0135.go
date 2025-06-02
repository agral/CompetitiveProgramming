package lc0135

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range n {
		candies[i] = 1
	}

	// Sweep left-to-right:
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Sweep right-to-left:
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// Return the sum of candies:
	ans := 0
	for i := range n {
		ans += candies[i]
	}
	return ans
}
