package lc2975

import "testing"

func Test_maximizeSquareArea(t *testing.T) {
	testcases := []struct {
		h        int
		w        int
		hFences  []int
		vFences  []int
		expected int
	}{
		{4, 3, []int{2, 3}, []int{2}, 4},
		{6, 7, []int{2}, []int{4}, -1},
	}

	for i, tc := range testcases {
		actual := maximizeSquareArea(tc.h, tc.w, tc.hFences, tc.vFences)
		if actual != tc.expected {
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
