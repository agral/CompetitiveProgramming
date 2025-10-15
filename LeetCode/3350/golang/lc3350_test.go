package lc3350

import "testing"

func Test_maxIncreasingSubarrays(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3},
		{[]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, 2},
	}

	for i, tc := range testcases {
		actual := maxIncreasingSubarrays(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
