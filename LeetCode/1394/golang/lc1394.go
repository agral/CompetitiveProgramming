package lc1394

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func findLucky(arr []int) int {
	frequencies := make(map[int]int)
	for _, val := range arr {
		frequencies[val] += 1
	}

	ans := -1
	for k, v := range frequencies {
		if k == v {
			ans = max(ans, k)
		}
	}
	return ans
}
