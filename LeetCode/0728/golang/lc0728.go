package lc0728

// Subjective level: easy+
// Solved on: 2026-01-25

// Runtime complexity: O(log10(num))
// Auxiliary space: O(1)
func isSelfDividing(num int) bool {
	n := num
	for n > 0 {
		d := n % 10
		if d == 0 {
			return false
		}
		if num%d != 0 {
			return false
		}
		n /= 10
	}
	return true
}

// Runtime complexity: O(k*log10(num)), where k == right - left + 1
// Auxiliary space: O(z); where z is the number of self-dividing numbers in that range.
// z is at most right-left+1 (for example for the range [1-9]). It can be less than that.
func selfDividingNumbers(left int, right int) []int {
	ans := []int{}
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			ans = append(ans, i)
		}
	}
	return ans
}
