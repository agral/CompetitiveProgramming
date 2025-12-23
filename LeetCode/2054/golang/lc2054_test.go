package lc2054

import "testing"

func Test_maxTwoEvents(t *testing.T) {
	testcases := []struct {
		events   [][]int
		expected int
	}{
		{[][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}, 4},
		{[][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}, 5},
		{[][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}, 8},
	}

	for i, tc := range testcases {
		actual := maxTwoEvents(tc.events)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.events, tc.expected, actual)
		}
	}
}

