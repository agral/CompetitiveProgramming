package lc2843

// Runtime complexity: O(high).
//
//	Note: It could be argued that since 1 <= high <= 10^4,
//	      it's actually O(10^4) == O(1) - but that's pushing it.
//
// Auxiliary space complexity: O(1)
func countSymmetricIntegers(low int, high int) int {
	ans := 0
	for num := low; num <= high; num++ {
		if isSymmetric(num) {
			ans++
		}
	}
	return ans
}

func isSymmetric(num int) bool {
	if num >= 10 && num <= 99 {
		return (num/10 == num%10)
	}
	if num >= 1000 && num <= 9999 {
		leftHalf, rightHalf := num/100, num%100
		/*
			leftDigitSum := leftHalf/10 + leftHalf%10
			rightDigitSum := rightHalf/10 + rightHalf%10
			return leftDigitSum == rightDigitSum
		*/
		return (leftHalf/10+leftHalf%10 == rightHalf/10+rightHalf%10)
	}
	return false
}
