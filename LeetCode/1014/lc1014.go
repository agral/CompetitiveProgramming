package lc1014

// Runtime complexity: O(n), the slice is iterated over just once.
// Space complexity: O(1); need constant space for two ints, ans and bestSeen.

func maxScoreSightseeingPair(values []int) int {
	ans := 0

	// holds the value of the best sightseeing place seen so far,
	// not including current place, taking the distance to it into account.
	bestSeen := 0

	for _, value := range values {
		ans = max(ans, value+bestSeen)
		bestSeen = max(bestSeen, value) - 1
	}
	return ans
}
