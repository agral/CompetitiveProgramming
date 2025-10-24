package lc2048

// Runtime complexity: O(1000000 log(1000000)) == O(1)
// Auxiliary space: O(1)
// Subjective level: a mid backtracking/counting problem. Did not enjoy it.
func nextBeautifulNumber(n int) int {
	// The following results in a WA. Of course the max number will be greater than that,
	// it's actually 1224444.
	//for k := n + 1; k < 1_000_001; k++ {
	for k := n + 1; k <= 1_224_444; k++ {
		if IsBalancedNumber(k) {
			return k
		}
	}
	return -42
}

func IsBalancedNumber(n int) bool {
	digitsCount := make([]int, 10)
	for n > 0 {
		if n%10 == 0 {
			return false
		}
		digitsCount[n%10] += 1
		n /= 10
	}
	for i := 1; i < 10; i++ {
		if digitsCount[i] > 0 && digitsCount[i] != i {
			return false
		}
	}
	return true
}
