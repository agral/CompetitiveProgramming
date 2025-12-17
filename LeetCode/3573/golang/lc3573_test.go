package lc3573

import "testing"

func Test_maximumProfit(t *testing.T) {
	testcases := []struct {
		prices   []int
		k        int
		expected int64
	}{
		{[]int{1, 7, 9, 8, 2}, 2, int64(14)},
		{[]int{12, 16, 19, 19, 8, 1, 19, 13, 9}, 3, int64(36)},
	}

	for i, tc := range testcases {
		actual := maximumProfit(tc.prices, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.prices, tc.k, tc.expected, actual)
		}
	}
}
