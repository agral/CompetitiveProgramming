package lc0808

// Runtime complexity: O(n), but can be argued that it's actually O(1) as n <= 4800.
// Auxiliary space: O(n^2), but can also be argued that it's actually O(192^2) == O(1).
func soupServings(n int) float64 {
	if n >= 4800 {
		return 1.0
	}
	const MEMO_SIZE int = 192 // == 4800 / 25
	memo := make([][]float64, MEMO_SIZE)
	for i := range MEMO_SIZE {
		memo[i] = make([]float64, MEMO_SIZE)
	}
	var dfs func(a, b int) float64
	dfs = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		} else if a <= 0 {
			return 1.0
		} else if b <= 0 {
			return 0.0
		}

		if memo[a][b] == 0.0 {
			memo[a][b] = 0.25 * (dfs(a-4, b) + dfs(a-3, b-1) + dfs(a-2, b-2) + dfs(a-1, b-3))
		}
		return memo[a][b]
	}
	return dfs((n+24)/25, (n+24)/25)
}
