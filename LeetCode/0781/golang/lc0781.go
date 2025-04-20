package lc0781

// Runtime complexity: O(n)
// Aux space complexity: O(1000) == O(1)
func numRabbits(answers []int) int {
	ans := 0
	minimums := make(map[int]int, 1000)
	for _, answer := range answers {
		if minimums[answer]%(answer+1) == 0 {
			ans += (answer + 1)
		}
		minimums[answer] += 1
	}
	return ans
}
