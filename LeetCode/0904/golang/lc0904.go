package lc0904

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func totalFruit(fruits []int) int {
	ans := 0
	count := make(map[int]int)
	l := 0
	for r := range fruits {
		count[fruits[r]] += 1
		for len(count) > 2 {
			count[fruits[l]] -= 1
			if count[fruits[l]] == 0 {
				delete(count, fruits[l])
			}
			l += 1
		}
		ans = max(ans, r-l+1)
	}
	return ans
}
