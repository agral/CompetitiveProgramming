package lc3105

import "testing"

func Test_longestMonotonicSubarray(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 4, 3, 3, 2}, 2}, // e.g. 1,4 increasing; 4,3 decreasing
		{[]int{3, 3, 3, 3}, 1},    // any [3] is the best increasing / decreasing subarray, can't do better.
		{[]int{3, 2, 1}, 3},       // [3, 2, 1] is best decreasing subarray.
	}

	for i, tc := range testcases {
		actual := longestMonotonicSubarray(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
