package lc3418

import "testing"

func Test_maximumAmount(t *testing.T) {
	testcases := []struct {
		coins    [][]int
		expected int
	}{
		{[][]int{
			{0, 1, -1},
			{1, -2, 3},
			{2, -3, 4},
		}, 8},
		{[][]int{
			{10, 10, 10},
			{10, 10, 10},
		}, 40},
	}

	for i, tc := range testcases {
		actual := maximumAmount(tc.coins)
		if actual != tc.expected {
			t.Errorf("Testcase maximumAmount#%02d (%v) failed: want %d, got %d",
				i+1, tc.coins, tc.expected, actual)
		}
	}
}
