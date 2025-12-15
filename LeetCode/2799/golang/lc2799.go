package lc2799

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium/hard
func countCompleteSubarrays(nums []int) int {
	const MAX int = 2000

	// Count the total number of distinct elements in nums:
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	fullDistinct := len(counter)

	left := 0
	currDistinct := 0
	currCount := make(map[int]int, MAX+1)
	ans := 0

	for _, num := range nums {
		if currCount[num] == 0 {
			currDistinct += 1
		}
		currCount[num] += 1

		for currDistinct == fullDistinct {
			currCount[nums[left]] -= 1
			if currCount[nums[left]] == 0 {
				currDistinct -= 1
			}
			left += 1
		}
		ans += left
	}

	return ans
}
