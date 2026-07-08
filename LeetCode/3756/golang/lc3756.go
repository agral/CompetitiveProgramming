package lc3756

const MAX = 100001
const MOD int64 = 1_000_000_007

// Runtime complexity: O(MAX)
// Auxiliary space: O(MAX)
func generatePowersOfTen() []int64 {
	powers := make([]int64, MAX)
	powers[0] = 1
	for i := 1; i < MAX; i++ {
		powers[i] = (10 * powers[i-1]) % MOD
	}
	return powers
}

var POWERS = generatePowersOfTen()

// Runtime complexity: O(n+q), where n is the len of the string s, and q is the number of queries.
// Auxiliary space: O(n).
// Note: not counting the complexity of `generatePowersOfTen()`, this gets resolved at compile time.
// Subjective level: hard- (i.e. an easier hard, but harder than a hard medium).
// Solved on: 2026-07-08
func sumAndMultiply(s string, queries [][]int) []int {
	L := len(s)
	sumDigits := make([]int, L+1)
	countNonzero := make([]int, L+1)
	prefixSum := make([]int64, L+1)

	for i := range L {
		digit64 := int64(s[i] - '0')
		sumDigits[i+1] = sumDigits[i] + int(digit64)
		if digit64 == 0 {
			prefixSum[i+1] = prefixSum[i]
			countNonzero[i+1] = countNonzero[i]
		} else {
			prefixSum[i+1] = (10*prefixSum[i] + digit64) % MOD
			countNonzero[i+1] = countNonzero[i] + 1
		}
	}

	ans := make([]int, len(queries))
	for idx, query := range queries {
		l, r := query[0], query[1]
		nonzero := countNonzero[r+1] - countNonzero[l]
		range_sum := int64(sumDigits[r+1] - sumDigits[l])

		// Literally following the challenge's description:
		// "Form a new integer x by concatenating all the non-zero digits from the substring
		// in their original order. If there are no non-zero digits, x = 0."
		x := (prefixSum[r+1] - prefixSum[l]*POWERS[nonzero]%MOD + MOD) % MOD
		ans[idx] = int(range_sum * x % MOD)
	}

	return ans
}
