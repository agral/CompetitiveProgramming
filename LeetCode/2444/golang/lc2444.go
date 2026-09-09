package lc2444

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func countSubarrays(nums []int, minK int, maxK int) int64 {
	ans := int64(0)
	lastIndexWrong, lastIndexMinK, lastIndexMaxK := -1, -1, -1

	for idx, num := range nums {
		if num < minK || num > maxK {
			lastIndexWrong = idx
		} else {
			if num == minK {
				lastIndexMinK = idx
			}
			// Note: not else-if, as minK and maxK might be the same number!
			if num == maxK {
				lastIndexMaxK = idx
			}
		}

		ans += int64(max(0, min(lastIndexMinK, lastIndexMaxK)-lastIndexWrong))
	}
	return ans
}
