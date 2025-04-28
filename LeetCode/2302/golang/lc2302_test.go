package lc2302

import "testing"

func Test_countSubarrays(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int64
		expected int64
	}{
		{[]int{2, 1, 4, 3, 5}, 10, 6},
		{[]int{1, 1, 1}, 5, 5},
	}

	for i, tc := range testcases {
		actual := countSubarrays(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v|%d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
