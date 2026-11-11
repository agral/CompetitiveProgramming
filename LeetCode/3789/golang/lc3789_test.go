package lc3789

import "testing"

func Test_minimumCost(t *testing.T) {
	testcases := []struct {
		cost1    int
		cost2    int
		costBoth int
		need1    int
		need2    int
		expected int64
	}{
		{3, 2, 1, 3, 2, int64(3)},
		{5, 4, 15, 2, 3, int64(22)},
		{5, 4, 15, 0, 0, int64(0)},
	}

	for i, tc := range testcases {
		actual := minimumCost(tc.cost1, tc.cost2, tc.costBoth, tc.need1, tc.need2)
		if actual != tc.expected {
			t.Errorf("Testcase minimumCost#%02d (cost1=%d, cost2=%d, costBoth=%d,"+
				"need1=%d, need2=%d) failed: want %d, got %d",
				i+1, tc.cost1, tc.cost2, tc.costBoth, tc.need1, tc.need2, tc.expected, actual)
		}
	}
}
