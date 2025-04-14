package lc1534

import "testing"

func Test_countGoodTriplets(t *testing.T) {
	testcases := []struct {
		arr      []int
		a        int
		b        int
		c        int
		expected int
	}{
		{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3, 4},
		{[]int{1, 1, 2, 2, 3}, 0, 0, 1, 0},
	}

	for i, tc := range testcases {
		actual := countGoodTriplets(tc.arr, tc.a, tc.b, tc.c)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.arr, tc.expected, actual)
		}
	}
}
