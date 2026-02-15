package lc0067

// Runtime complexity: O(n), where n is the size of the largest of the two binary strings.
// Auxiliary space: O(n) (for the answer which has to be of size up to n+1
// Subjective level: easy
// Solved on: 2026-02-15
func addBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	// 1. Make sure that a is the longer of the two strings.
	// Switch `a` and `b` if it's not the case:
	if lenB > lenA {
		a, b = b, a
		lenA, lenB = lenB, lenA
	}

	// 2. Pre-allocate a slice of length lenA to hold the result.
	// Note: lenA+1 is the maximum length of the binary string resulting
	// from adding the two numbers. But we'll care about nonzero carry at the end.
	result := make([]byte, lenA)

	// 3. Perform the long addition of binary digits of both a and b:
	carry := 0
	j := lenB - 1
	for i := lenA - 1; i >= 0; i-- {
		// lhs and rhs are either 0 or 1, read from the input strings:
		lhs, rhs := 0, 0
		if a[i] == '1' {
			lhs = 1
		}
		if j >= 0 {
			if b[j] == '1' {
				rhs = 1
			}
		} // and when j < 0, rhs stays 0 - there's no more digits in b to use.
		sum := lhs + rhs + carry // each operand is either 0 or 1, so sum is anything between 0 and 3.
		result[i] = '0'
		if sum == 1 || sum == 3 { // equivalent to `if sum % 2 == 1`, but slightly faster.
			result[i] = '1'
		}
		carry = sum / 2 // for sum == 2 or sum == 3, carry is 1; for 0 or 1 it remains 0.
		j -= 1
		//fmt.Printf("i=%d, j=%d, r=%d, carry=%d\n")
	}

	if carry == 1 {
		return "1" + string(result)
	}
	return string(result)
}
