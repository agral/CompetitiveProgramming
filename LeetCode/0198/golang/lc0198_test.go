package lc0198

import "testing"

func Test_rob(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},     // rob houses #0(1) and #2(3), the total is 4
		{[]int{2, 7, 9, 3, 1}, 12}, //rob houses #0, #2 and #4 for the total of 12
		{[]int{42}, 42},            // a degenerated case: only one house along the street.
	}

	for i, tc := range testcases {
		actual := rob(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
