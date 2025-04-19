package lc2563

import "testing"

func Test_countFairPairs(t *testing.T) {
	testcases := []struct {
		nums     []int
		lower    int
		upper    int
		expected int64
	}{
		{[]int{0, 1, 7, 4, 4, 5}, 3, 6, 6},
		{[]int{1, 7, 9, 2, 5}, 11, 11, 1},
	}

	for i, tc := range testcases {
		actual := countFairPairs(tc.nums, tc.lower, tc.upper)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v, %d, %d) failed: want %d, got %d",
				i+1, tc.nums, tc.lower, tc.upper, tc.expected, actual)
		}
	}
}
