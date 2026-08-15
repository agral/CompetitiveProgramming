package lc3702

import "testing"

func Test_longestSubsequence(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{2, 3, 4}, 3},
	}

	for i, tc := range testcases {
		actual := longestSubsequence(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase longestSubsequence#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
