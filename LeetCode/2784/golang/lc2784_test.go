package lc2784

import "testing"

func Test_isGood(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 1, 3}, false},
		{[]int{1, 3, 3, 2}, true},
		{[]int{1, 1}, true},
		{[]int{3, 4, 4, 1, 2, 1}, false}, // extra "1".
	}

	for i, tc := range testcases {
		actual := isGood(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase isGood#%02d (%v) failed: want %t, got %t",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
