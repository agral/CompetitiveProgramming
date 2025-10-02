package lc3100

// Runtime complexity: O(log(n))
// Auxiliary space: O(1)
// Subjective level: easy
func maxBottlesDrunk(numBottles int, numExchange int) int {
	// This is almost the same as 1518/Water Bottles; the only difference being that
	// numExchange gets increased by one after each exchange.
	ans := 0
	full := numBottles
	empty := 0
	for {
		// drink:
		empty += full
		ans += full
		full = 0

		// exchange empty bottles:
		for empty >= numExchange {
			empty -= numExchange
			full += 1
			numExchange += 1
		}

		// if no bottles exchanged in previous step, no further progress can be done:
		if full == 0 {
			return ans
		}
	}
}
