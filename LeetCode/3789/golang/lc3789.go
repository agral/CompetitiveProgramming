package lc3789

// Runtime complexity: O(1)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-11-11
func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	ans := int64(0)

	// Step 1: convert all data to int64, as stuff will probably overflow (it's got like 50% acceptance rate):
	c1 := int64(cost1)
	c2 := int64(cost2)
	cB := int64(costBoth)
	n1 := int64(need1)
	n2 := int64(need2)

	// Step 2: keep buying either cost1+cost2 or costBoth, whichever is cheaper,
	// until min(need1, need2) is reached.
	minNeeded := min(n1, n2)
	ans += min(c1+c2, cB) * minNeeded

	// Step 3: buy the remaining need1-minNeeded / need2-minNeeded needed items.
	// Note that it might be more optimal to buy at costBoth than at cost1/cost2.
	// Also note that Step 1 might be enough (for cases where need1 == need2).
	if n1 > minNeeded {
		ans += min(c1, cB) * (n1 - minNeeded)
	}
	if n2 > minNeeded {
		ans += min(c2, cB) * (n2 - minNeeded)
	}

	// After Step 3 we're done, return the answer.
	return ans
}
