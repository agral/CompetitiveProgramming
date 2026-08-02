package lc0877

import "testing"

func Test_stoneGame(t *testing.T) {
	testcases := []struct {
		piles    []int
		expected bool
	}{
		{[]int{5, 3, 4, 5}, true},
		{[]int{3, 7, 2, 3}, true},
	}

	for i, tc := range testcases {
		actual := stoneGame(tc.piles)
		if actual != tc.expected {
			t.Errorf("Testcase stoneGame#%02d (%v) failed: want %t, got %t",
				i+1, tc.piles, tc.expected, actual)
		}
	}
}
