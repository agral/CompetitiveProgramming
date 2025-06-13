package lc2616

import "testing"

func Test_minimizeMax(t *testing.T) {
	testcases := []struct {
		nums     []int
		p        int
		expected int
	}{
		{[]int{10, 1, 2, 7, 1, 3}, 2, 1}, // choose (1, 1) and (2, 3); max diff is |2-3| == 1.
		{[]int{4, 2, 1, 2}, 1, 0},        // choose (2, 2); max diff is 0.
	}

	for i, tc := range testcases {
		actual := minimizeMax(tc.nums, tc.p)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.p, tc.expected, actual)
		}
	}
}
