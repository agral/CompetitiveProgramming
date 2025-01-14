package lc2657

func findThePrefixCommonArray(A []int, B []int) []int {
	SZ := len(A)
	ans := make([]int, SZ)
	seen := make([]int, SZ+1)
	commonPrefixCounter := 0

	for i := 0; i < SZ; i++ {
		seen[A[i]]++
		if seen[A[i]] == 2 {
			commonPrefixCounter++
		}
		seen[B[i]]++
		if seen[B[i]] == 2 {
			commonPrefixCounter++
		}
		ans[i] = commonPrefixCounter
	}

	return ans
}
