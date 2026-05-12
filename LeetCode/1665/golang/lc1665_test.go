package lc1665

import "testing"

func Test_minimumEffort(t *testing.T) {
	testcases := []struct {
		tasks    [][]int
		expected int
	}{
		{[][]int{{1, 2}, {2, 4}, {4, 8}}, 8},
		{[][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}}, 32},
	}

	for i, tc := range testcases {
		actual := minimumEffort(tc.tasks)
		if actual != tc.expected {
			t.Errorf("Testcase minimumEffort#%02d (%v) failed: want %d, got %d",
				i+1, tc.tasks, tc.expected, actual)
		}
	}
}
