package lc3653

// Runtime complexity: O(len(queries) * len(nums))
// Auxiliary space: O(nums)
// Subjective level: medium-
// Solved on: 2026-04-08
func xorAfterQueries(nums []int, queries [][]int) int {
	const MOD int64 = 1_000_000_007

	// 2. Process all the queries:
	for _, query := range queries {
		for i := query[0]; i <= query[1]; i += query[2] {
			// avoid int32 overflow:
			nums[i] = int((int64(query[3]) * int64(nums[i])) % MOD)
		}
	}

	// 3. Calculate the XOR of all the elements:
	ans := 0
	for _, num := range nums {
		ans ^= num
	}

	// 4. Return the answer:
	return ans
}
