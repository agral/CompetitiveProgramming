package lc3392

import "testing"

func Test_countSubarrays(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 1, 4, 1}, 1},
		{[]int{1, 1, 1}, 0},
	}

	for i, tc := range testcases {
		actual := countSubarrays(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
