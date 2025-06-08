package lc0386

// Runtime complexity: O(n) [as required by the problem statement]
// Auxiliary space: O(1) [as required by the problem statement]
func lexicalOrder(n int) []int {
	ans := make([]int, n)

	a := 0
	curr := 1
	for a < n {
		ans[a] = curr
		a++
		if 10*curr <= n {
			curr *= 10
		} else {
			for curr%10 == 9 || curr == n {
				curr = curr / 10
			}
			curr += 1
		}
	}

	return ans
}
