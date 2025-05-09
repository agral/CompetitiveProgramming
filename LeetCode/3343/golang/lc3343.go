package lc3343

// Runtime complexity:
// Auxiliary space:
func countBalancedPermutations(num string) int {
	// Count the sum of all the digits:
	sum := 0
	for _, d := range num {
		digit := int(d - '0')
		sum += digit
	}
	if sum%2 == 1 {
		// if the sum of all digits is odd, the numbers can not be divided into
		// two sets of same sums.
		return 0
	}
	ans := -1
	return ans
}

func factorial(num int) int64 {
	if num <= 1 {
		return 1
	}
	ans := int64(1)
	for k := int64(2); k <= int64(num); k++ {
		ans *= k
	}
	return ans
}
