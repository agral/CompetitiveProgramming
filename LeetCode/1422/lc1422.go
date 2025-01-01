package lc1422

func maxScore(s string) int {
	sz := len(s)
	// The string gets divided at some index i into left s[:i] and right s[i:] substrings.
	// E.g.: "00_111": i=2; left="00", right="111"
	// index: 01_234

	// holds the total number of zeroes in the left substring;
	// i.e. "How much zeroes are there before me?" (not including this index).
	// Because both left and right substrings can not be empty, index 0 will not be used
	// (as such division allocates an empty string to the left and the entire input string to the right part).
	leftSumZeroes := make([]int, sz)
	for i := 1; i < sz; i++ {
		leftSumZeroes[i] = leftSumZeroes[i-1]
		if s[i-1] == '0' {
			leftSumZeroes[i] += 1
		}
	}

	// holds the total number of ones in the right substring;
	// ie.e "How much ones are there after me?" (including this index too).
	// rightSumOnes[i] = sum("1", s[i+1:])
	rightSumOnes := make([]int, sz)
	for i := sz - 1; i > 0; i-- {
		if i < sz-1 {
			rightSumOnes[i] = rightSumOnes[i+1]
		}
		if s[i] == '1' {
			rightSumOnes[i] += 1
		}
	}

	ans := 0
	for i := 1; i < sz; i++ {
		ans = max(ans, leftSumZeroes[i]+rightSumOnes[i])
	}

	return ans
}
