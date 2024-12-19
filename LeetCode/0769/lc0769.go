package lc0769

// Runtime complexity: O(n).
// Space complexity: O(1).
// Tags: Array, Greedy

func maxChunksToSorted(arr []int) int {
	ans := 0
	max_seen := -1
	for i := range arr {
		max_seen = max(max_seen, arr[i])
		if max_seen == i {
			ans += 1
		}
	}
	return ans
}
