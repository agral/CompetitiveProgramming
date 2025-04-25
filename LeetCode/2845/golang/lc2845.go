package lc2845

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	ans := int64(0)
	prefix := 0
	prefixCount := make(map[int]int)
	prefixCount[0] = 1

	for _, num := range nums {
		if num%modulo == k {
			prefix = (prefix + 1) % modulo
		}
		ans += int64(prefixCount[(prefix-k+modulo)%modulo])
		prefixCount[prefix] += 1
	}

	return ans
}
