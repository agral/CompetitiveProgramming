package lc0038

import (
	"strconv"
)

// Runtime complexity: O(len(ans))
// Space complexity: O(len(ans))
func countAndSay(n int) string {
	ans := "1"
	for range n - 1 {
		next := ""
		i := 0
		for i < len(ans) {
			count := 1
			for i+1 < len(ans) && ans[i] == ans[i+1] {
				count++
				i++
			}
			next += strconv.Itoa(count) + string(ans[i])
			i++
		}
		ans = next
	}
	return ans
}
