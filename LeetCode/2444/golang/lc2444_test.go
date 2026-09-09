package lc2444

import "testing"

func Test_countSubarrays(t *testing.T) {
	testcases := []struct {
		nums     []int
		minK     int
		maxK     int
		expected int64
	}{
		{[]int{1, 3, 5, 2, 7, 5}, 1, 5, int64(2)},
		{[]int{1, 1, 1, 1}, 1, 1, int64(10)},
	}

	for i, tc := range testcases {
		actual := countSubarrays(tc.nums, tc.minK, tc.maxK)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v|minK=%d, maxK=%d) failed: want %d, got %d",
				i+1, tc.nums, tc.minK, tc.maxK, tc.expected, actual)
		}
	}
}
