package lc2264

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func largestGoodInteger(num string) string {
	ans := ""
	for i := 2; i < len(num); i++ {
		if num[i-2] == num[i-1] && num[i-1] == num[i] {
			ans = max(ans, num[i-2:i+1])
		}
	}
	return ans
}
