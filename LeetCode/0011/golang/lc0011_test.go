package lc0011

import "testing"

func Test_maxArea(t *testing.T) {
	testcases := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for i, tc := range testcases {
		actual := maxArea(tc.height)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.height, tc.expected, actual)
		}
	}
}
