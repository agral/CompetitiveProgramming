package lc3719

// Runtime complexity: O(n^2)
// Auxiliary space: O(n)
// Subjective level: medium. Needs a different approach than the brute force recommended
// > in the hints. So the hints are misleading for this one. Seems that quite a lot
// > of recently added LC questions have this exact problem.
//
// Solved on: 2026-02-10
// Beats 93.33% of the solutions submitted until 2026-02-10, 10:34 CET.
func longestBalanced(nums []int) int {
	LEN := len(nums)
	ans := 0
	// 1. Generate every possible subarray that has a chance of containing
	// one even & one odd number (i.e. all subarrays of length >= 2):
	for l := LEN; l >= 2; l-- {
		// 2. Consider a subarray starting at index 0, and of length l.
		//    keep track of the seen numbers & their count, and keep track
		//    of its stats - number of unique even & odd numbers:
		seen := make(map[int]int)
		numUniqueEven, numUniqueOdd := 0, 0
		for i := 0; i < l; i++ {
			seen[nums[i]] += 1
			if seen[nums[i]] == 1 {
				if nums[i]%2 == 0 {
					numUniqueEven += 1
				} else {
					numUniqueOdd += 1
				}
			}
		}
		if numUniqueEven == numUniqueOdd {
			return l
		}

		// 3. Now consider all the remaining subarrays of length l.
		// These have the same stats as prevous one, except with the
		// beginning number removed and the last number added.
		// Update the stats accordingly for each of these subarrays.
		for b := 1; b+l <= LEN; b++ {
			// nums[b-1] gets removed from the subarray:
			seen[nums[b-1]] -= 1
			if seen[nums[b-1]] == 0 {
				if nums[b-1]%2 == 0 {
					numUniqueEven -= 1
				} else {
					numUniqueOdd -= 1
				}
			}

			// nums[b+l-1] gets added to the subarray:
			seen[nums[b+l-1]] += 1
			if seen[nums[b+l-1]] == 1 {
				if nums[b+l-1]%2 == 0 {
					numUniqueEven += 1
				} else {
					numUniqueOdd += 1
				}
			}

			// If the count of unique even & odd numbers is the same,
			// this is the longest balanced subarray - return its length `l`.
			if numUniqueEven == numUniqueOdd {
				return l
			}
		}
	}

	return ans
}
