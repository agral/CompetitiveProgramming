package lc3349

import "testing"

func Test_hasIncreasingSubarrays(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3, true},
		{[]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, 5, false},
	}

	for i, tc := range testcases {
		actual := hasIncreasingSubarrays(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.nums, tc.expected, actual)
		}
	}
}
