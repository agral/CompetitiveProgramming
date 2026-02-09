package lc1475

func finalPrices(prices []int) []int {
	ans := prices[:]
	for i, price := range prices {
		for _, p := range prices[i+1:] {
			if p <= price {
				ans[i] -= p
				break // inner for loop
			}
		}
	}
	return ans
}
