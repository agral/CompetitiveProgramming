package lc2540

// Runtime complexity: O(n), where n == len(nums1) + len(nums2) - both worst and average case.
// Auxiliary space: O(1)
// Subjective level: easy peasy
// Solved on: 2026-05-19
func getCommon(nums1 []int, nums2 []int) int {
	pos1, pos2 := 0, 0
	L1, L2 := len(nums1), len(nums2)
	for pos1 < L1 && pos2 < L2 {
		if nums1[pos1] == nums2[pos2] {
			return nums1[pos1]
		} else if nums1[pos1] < nums2[pos2] {
			pos1 += 1
		} else {
			pos2 += 1
		}
	}
	return -1
}
