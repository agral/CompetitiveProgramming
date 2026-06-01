package lc2144

import "testing"

func Test_minimumCost(t *testing.T) {
	testcases := []struct {
		cost     []int
		expected int
	}{
		{[]int{1, 2, 3}, 5},
		{[]int{6, 5, 7, 9, 2, 2}, 23},
		{[]int{5, 5}, 10},
	}

	for i, tc := range testcases {
		actual := minimumCost(tc.cost)
		if actual != tc.expected {
			t.Errorf("Testcase minimumCost#%02d (%v) failed: want %d, got %d",
				i+1, tc.cost, tc.expected, actual)
		}
	}
}
