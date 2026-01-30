package lc0223

import "testing"

func Test_computeArea(t *testing.T) {
	testcases := []struct {
		ax1      int
		ay1      int
		ax2      int
		ay2      int
		bx1      int
		by1      int
		bx2      int
		by2      int
		expected int
	}{
		{-3, 0, 3, 4, 0, -1, 9, 2, 45},
		{-2, -2, 2, 2, -2, -2, 2, 2, 16},
	}

	for i, tc := range testcases {
		actual := computeArea(tc.ax1, tc.ay1, tc.ax2, tc.ay2, tc.bx1, tc.by1, tc.bx2, tc.by2)
		if actual != tc.expected {
			t.Errorf("Testcase computeArea#%02d failed: want %d, got %d",
				i+1, tc.expected, actual)
		}
	}
}
