package lc0611

import "slices"

// Assumes first <= second <= third.
func canFormTriangle(first int, second int, third int) bool {
	return first+second > third
}

// Runtime complexity: O(n^3)
// Auxiliary space: O(1)
// Subjective level: medium/easy.
func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	ans := 0
	slices.Sort(nums)
	for first := 0; first < len(nums)-2; first++ {
		for second := first + 1; second < len(nums)-1; second++ {
			for third := second + 1; third < len(nums); third++ {
				if canFormTriangle(nums[first], nums[second], nums[third]) {
					ans += 1
				}
			}
		}
	}

	return ans
}
