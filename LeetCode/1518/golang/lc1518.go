package lc1518

// Runtime complexity: O(log(n))
// Auxiliary space: O(1)
// Subjective level: easy
func numWaterBottles(numBottles int, numExchange int) int {
	ans := 0
	full := numBottles
	empty := 0
	for {
		// drink:
		empty += full
		ans += full
		full = 0

		// exchange empty bottles:
		setsExchanged := empty / numExchange
		empty -= setsExchanged * numExchange
		full += setsExchanged

		// if no bottles exchanged in previous step, no further progress can be done:
		if full == 0 {
			return ans
		}
	}
}
