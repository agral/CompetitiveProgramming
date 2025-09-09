package lc2327

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium+
func peopleAwareOfSecret(n int, delay int, forget int) int {
	const MOD = int64(1_000_000_007)
	// peopleWhoKnow[i] holds the count of people who know the secret on day i.
	peopleWhoKnow := make([]int64, n)
	peopleWhoKnow[0] = 1
	countKnowing := int64(0)
	for i := 1; i < n; i++ {
		if i-forget >= 0 {
			countKnowing -= int64(peopleWhoKnow[i-forget])
		}
		if i-delay >= 0 {
			countKnowing += int64(peopleWhoKnow[i-delay])
		}
		countKnowing = (countKnowing + MOD) % MOD
		peopleWhoKnow[i] = countKnowing
	}
	ans := int64(0)
	for i := n - forget; i < n; i++ {
		ans += peopleWhoKnow[i]
	}
	return int(ans % MOD)
}
