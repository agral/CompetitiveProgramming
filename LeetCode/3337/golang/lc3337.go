package lc3337

const MOD = 1_000_000_007
const MOD64 = int64(1_000_000_007)

func makeIdentityMatrix(size int) [][]int {
	ans := make([][]int, size)
	for i := range size {
		ans[i] = make([]int, size)
		ans[i][i] = 1
	}
	return ans
}

func makeTransformationMatrix(nums []int) [][]int {
	ans := make([][]int, 26)
	for i := range 26 {
		ans[i] = make([]int, 26)
	}
	for i := range nums {
		for step := 1; step <= nums[i]; step++ {
			ans[i][(i+step)%26] += 1
		}
	}

	return ans
}

// Returns the result of multiplying the matrices lhs and rhs together.
func multiply(lhs [][]int, rhs [][]int) [][]int {
	size := len(lhs)
	ans := make([][]int, size)
	for i := range size {
		ans[i] = make([]int, size)
	}

	for i := range size {
		for j := range size {
			for k := range size {
				ans[i][j] = int((int64(lhs[i][k])*int64(rhs[k][j]) + int64(ans[i][j])) % MOD64)
			}
		}
	}
	return ans
}

// Returns the result of raising the matrix m to the power n. Recursive.
func power(m [][]int, n int) [][]int {
	if n == 0 {
		return makeIdentityMatrix(len(m))
	}
	if n%2 == 0 {
		return power(multiply(m, m), n/2)
	}
	return multiply(m, power(m, n-1))
}

// Pretty similar to lc3335. But with one extra dimension :-)
// Runtime complexity: O(len(s) + t).
// Auxiliary space: O(26*26) == O(1).
func lengthAfterTransformations(s string, t int, nums []int) int {
	tm := makeTransformationMatrix(nums)
	tmRaised := power(tm, t)

	// lettersCount[i]: holds the initial count of individual letters.
	lettersCount := make([]int, 26)
	bytestring := []byte(s)
	for _, ch := range bytestring {
		lettersCount[int(ch-'a')]++
	}

	// totals[i]: the total count of 'a'+i after t transformations.
	totals := make([]int64, 26)

	for i := range 26 {
		for j := range 26 {
			totals[j] = (totals[j] + int64(lettersCount[i])*int64(tmRaised[i][j])) % MOD64
		}
	}

	ans := 0
	for k := range totals {
		// Note: totals[k] guaranteed to not overflow int; as it is always %MOD64'ed.
		ans = (ans + int(totals[k])) % MOD
	}
	return ans
}
