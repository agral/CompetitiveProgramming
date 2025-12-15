package lc2110

import "testing"

func Test_getDescentPeriods(t *testing.T) {
	testcases := []struct {
		prices   []int
		expected int64
	}{
		{[]int{3, 2, 1, 4}, int64(7)},
		{[]int{8, 6, 7, 7}, int64(4)},
		{[]int{1}, int64(1)},
	}

	for i, tc := range testcases {
		actual := getDescentPeriods(tc.prices)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.prices, tc.expected, actual)
		}
	}
}
