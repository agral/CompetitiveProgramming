package lc2364

import "testing"

func Test_countBadPairs(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int64
	}{
		{[]int{4, 1, 3, 3}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}

	for i, tc := range testcases {
		actual := countBadPairs(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
