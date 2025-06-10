package lc0875

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	testcases := []struct {
		piles    []int
		h        int
		expected int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
	}

	for i, tc := range testcases {
		actual := minEatingSpeed(tc.piles, tc.h)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.piles, tc.h, tc.expected, actual)
		}
	}
}
