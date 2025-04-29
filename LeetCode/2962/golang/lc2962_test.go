package lc2962

import "testing"

func Test_countSubarrays(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int64
	}{
		{[]int{1, 3, 2, 3, 3}, 2, 6},
		{[]int{1, 4, 2, 1}, 3, 0},
	}

	for i, tc := range testcases {
		actual := countSubarrays(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
