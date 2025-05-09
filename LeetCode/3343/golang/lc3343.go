package lc3343

const MOD = 1_000_000_007
const MOD64 = int64(MOD)

// Runtime complexity:
// Auxiliary space:
func countBalancedPermutations(num string) int {
	// Count the sum of all the digits, and the frequency of individual digits:
	sum := 0
	digits := getDigits(num)
	for digit, count := range digits {
		sum += digit * count
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

func getDigits(num string) []int {
	ans := make([]int, 10)
	for _, d := range num {
		ans[int(d-'0')]++
	}
	return ans
}

func calcFreq(digits []int) {
}

func calcNumPermutations(digits []int) int64 {
	ans := int64(1)
	return ans // not implemented so far
}
