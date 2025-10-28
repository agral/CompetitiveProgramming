package lc3354

import "testing"

func Test_countValidSelections(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 0, 2, 0, 3}, 2},
		{[]int{2, 3, 4, 0, 4, 1, 0}, 0},
	}

	for i, tc := range testcases {
		actual := countValidSelections(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
