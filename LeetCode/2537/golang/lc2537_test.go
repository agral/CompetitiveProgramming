package lc2537

import "testing"

func Test_countGood(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int64
	}{
		{[]int{1, 1, 1, 1, 1}, 10, 1},
		{[]int{3, 1, 4, 3, 2, 2, 4}, 2, 4},
	}

	for i, tc := range testcases {
		actual := countGood(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
