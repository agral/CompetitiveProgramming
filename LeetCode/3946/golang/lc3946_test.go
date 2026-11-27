package lc3946

import "testing"

func Test_maximumSaleItems(t *testing.T) {
	testcases := []struct {
		items    [][]int
		budget   int
		expected int
	}{
		{[][]int{{6, 2}, {2, 6}, {3, 4}}, 9, 4},
		{[][]int{{2, 4}, {3, 2}, {4, 1}, {6, 4}, {12, 4}}, 8, 10},
		{[][]int{{1, 1}, {1, 2}}, 1, 2},
	}

	for i, tc := range testcases {
		actual := maximumSaleItems(tc.items, tc.budget)
		if actual != tc.expected {
			t.Errorf("Testcase maximumSaleItems#%02d (%v | %d) failed: want %d, got %d",
				i+1, tc.items, tc.budget, tc.expected, actual)
		}
	}
}
