package lc3151

import "testing"

func Test_isArraySpecial(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1}, true},
		{[]int{2, 1, 4}, true},
		{[]int{4, 3, 1, 6}, false},
	}

	for i, tc := range testcases {
		actual := isArraySpecial(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.nums, tc.expected, actual)
		}
	}
}
