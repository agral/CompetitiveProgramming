package lc3069

import "slices"

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: trivially easy
// Solved on: 2026-08-20
func resultArray(nums []int) []int {
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	len1 := 0
	len2 := 0
	for i := 2; i < len(nums); i++ {
		if arr1[len1] > arr2[len2] {
			// append ith num to arr1.
			arr1 = append(arr1, nums[i])
			len1++
		} else {
			// append ith num to arr2.
			arr2 = append(arr2, nums[i])
			len2++
		}
	}
	return slices.Concat(arr1, arr2)
}
