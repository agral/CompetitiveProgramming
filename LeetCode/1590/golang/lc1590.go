package lc1590

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium/hard
// Solved on: 2025-11-30
func minSubarray(nums []int, p int) int {
	// WA #01: it's not allowed to remove the entire array!
	// so can't use an actual BIG_NUM, or rather when I tried this approach
	// the code could not register that the entire array is being deleted.
	//
	// const BIG_NUM = math.MaxInt32

	BIG_NUM := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	rem := sum % p
	if rem == 0 {
		return 0
	}

	prefixToIndex := map[int]int{}
	prefixToIndex[0] = -1
	prefix := 0
	ans := BIG_NUM

	for i, num := range nums {
		prefix += num
		prefix %= p
		target := (prefix - rem + p) % p
		val, keyExists := prefixToIndex[target]
		if keyExists {
			ans = min(ans, i-val)
		}
		prefixToIndex[prefix] = i
	}

	if ans == BIG_NUM {
		return -1
	}

	return ans
}
