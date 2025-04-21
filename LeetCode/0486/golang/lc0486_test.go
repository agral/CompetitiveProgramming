package lc0486

import "testing"

func Test_predictTheWinner(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 5, 2}, false},
		{[]int{1, 5, 233, 7}, true},
	}

	for i, tc := range testcases {
		actual := predictTheWinner(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.nums, tc.expected, actual)
		}
	}
}
