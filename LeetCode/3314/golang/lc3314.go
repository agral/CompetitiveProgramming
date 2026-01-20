package lc3314

// Runtime complexity: O(N * sum(log(nums[i]));
// - hard to estimate, but higher than N, and surely lower than N^2.
// - probably close to O(N*logN) or maybe to O(N).
// Auxiliary space: O(N) (for holding the answer; otherwise O(1)).
// Subjective level: easy+. Requires bit shifting, but nothing fancy.
// Solved on: 2026-01-20
func minBitwiseArray(nums []int) []int {
	N := len(nums)
	ans := make([]int, N)

	getFirstSetBitOfLastGroupOfSetBits := func(num int) int {
		first := 1
		for num&first > 0 {
			first <<= 1
		} // note: this overshoots, i.e. actually points to the zero just before the actual answer.
		return first >> 1
	}

	for i, num := range nums {
		if num == 2 {
			ans[i] = -1
		} else {
			ans[i] = num - getFirstSetBitOfLastGroupOfSetBits(num)
		}
	}

	return ans
}
