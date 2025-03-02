package lc2570

// Runtime complexity: O(len(nums1)+len(nums2)). Optimal.
// Auxiliary space complexity: O(1001) == O(1); this may be suboptimal but is quite elegant.
func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	const MAX int = 1000 // 1 <= id_i, val_i <= 1000.
	entries := make([]int, MAX+1)
	for _, pair := range nums1 {
		// pair has the ID at index 0, and value at index 1.
		entries[pair[0]] += pair[1]
	}
	for _, pair := range nums2 {
		entries[pair[0]] += pair[1]
	}

	// Construct the answer array from nonzero entries in `entries` array:
	var ans [][]int
	for id, value := range entries {
		if value != 0 {
			ans = append(ans, []int{id, value})
		}
	}
	return ans
}
