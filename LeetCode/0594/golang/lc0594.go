package lc0594

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func findLHS(nums []int) int {
	freq := map[int]int{}
	for _, val := range nums {
		freq[val] += 1
	}

	ans := 0
	for num, count := range freq {
		if freq[num+1] > 0 {
			ans = max(ans, count+freq[num+1])
		}
	}
	return ans
}
