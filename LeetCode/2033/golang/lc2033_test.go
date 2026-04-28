package lc2033

import "testing"

func Test_minOperations(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		x        int
		expected int
	}{
		{[][]int{{2, 4}, {6, 8}}, 2, 4},
		{[][]int{{1, 5}, {2, 3}}, 1, 5},
		{[][]int{{1, 2}, {3, 4}}, 2, -1},
	}

	for i, tc := range testcases {
		actual := minOperations(tc.grid, tc.x)
		if actual != tc.expected {
			t.Errorf("Testcase minOperations#%02d (%v | %d) failed: want %d, got %d",
				i+1, tc.grid, tc.x, tc.expected, actual)
		}
	}
}
