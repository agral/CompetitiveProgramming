package lc1304

// Runtime complexity: O(n)
// Auxiliary space: O(n) (necessary to provide the answer in form of a slice; other than that - O(1).
// Subjective level: easy.
func sumZero(n int) []int {
	ans := make([]int, n)
	halfRange := n / 2
	for i := 1; i <= halfRange; i++ {
		ans[2*(i-1)] = -i
		ans[2*(i-1)+1] = i
	}
	if n%2 == 1 {
		ans[n-1] = 0
	}
	return ans
}
