package lc2845

import "testing"

func Test_countInterestingSubarrays(t *testing.T) {
	testcases := []struct {
		nums     []int
		modulo   int
		k        int
		expected int64
	}{
		{[]int{3, 2, 4}, 2, 1, int64(3)},
		{[]int{3, 1, 9, 6}, 3, 0, int64(2)},
	}

	for i, tc := range testcases {
		actual := countInterestingSubarrays(tc.nums, tc.modulo, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d, %d) failed: want %d, got %d",
				i+1, tc.nums, tc.modulo, tc.k, tc.expected, actual)
		}
	}
}
