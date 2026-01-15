package lc2943

import "testing"

func Test_maximizeSquareHoleArea(t *testing.T) {
	testcases := []struct {
		n        int
		m        int
		hBars    []int
		vBars    []int
		expected int
	}{
		{2, 1, []int{2, 3}, []int{2}, 4},
		{1, 1, []int{2}, []int{2}, 4},
		{2, 3, []int{2, 3}, []int{2, 4}, 4},
	}

	for i, tc := range testcases {
		actual := maximizeSquareHoleArea(tc.n, tc.m, tc.hBars, tc.vBars)
		if actual != tc.expected {
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
