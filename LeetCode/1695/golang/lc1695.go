package lc1695

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func maximumUniqueSubarray(nums []int) int {
	// Two pointers + uniques' set:
	currValue := 0
	ans := 0
	seen := make(map[int]bool) // no sets in Golang's stdlib as far as I know.
	left := 0
	for right := range nums {
		// Remove the duplicate values from the left:
		for seen[nums[right]] {
			currValue -= nums[left]
			seen[nums[left]] = false
			left++
		}

		// Add this new number from the right, update the score:
		seen[nums[right]] = true
		currValue += nums[right]
		ans = max(ans, currValue)
	}
	return ans
}
