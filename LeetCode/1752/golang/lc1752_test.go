package lc1752

import "testing"

func Test_check(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{3, 4, 5, 1, 2}, true},
		{[]int{2, 1, 3, 4}, false},
		{[]int{1, 2, 3}, true},
	}

	for i, tc := range testcases {
		actual := check(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.nums, tc.expected, actual)
		}
	}
}
