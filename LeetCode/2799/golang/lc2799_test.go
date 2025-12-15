package lc2799

import "testing"

func Test_countCompleteSubarrays(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 1, 2, 2}, 4},
		{[]int{5, 5, 5, 5}, 10},
	}

	for i, tc := range testcases {
		actual := countCompleteSubarrays(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
