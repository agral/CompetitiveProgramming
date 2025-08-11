package lc2438

// Runtime complexity: O(len(queries))
// Auxiliary space: O(len(queries))
func productQueries(n int, queries [][]int) []int {
	const MOD = int64(1_000_000_007)
	const MAX_SET_BIT int = 30

	powers := []int64{}
	for i := range MAX_SET_BIT {
		if (n>>i)&1 == 1 {
			powers = append(powers, int64(1<<i))
		}
	}

	len_queries := len(queries)
	ans := make([]int, len_queries)
	for q := range queries {
		left := queries[q][0]
		right := queries[q][1]
		product := int64(1)
		for i := left; i <= right; i++ {
			product *= powers[i]
			product %= MOD
		}
		ans[q] = int(product)
	}

	return ans
}
