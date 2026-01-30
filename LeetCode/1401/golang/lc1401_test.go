package lc1401

import "testing"

func Test_checkOverlap(t *testing.T) {
	testcases := []struct {
		radius   int
		x        int
		y        int
		x1       int
		y1       int
		x2       int
		y2       int
		expected bool
	}{
		{1, 0, 0, 1, -1, 3, 1, true},
		{1, 1, 1, 1, -3, 2, -1, false},
		{1, 0, 0, -1, 0, 0, 1, true},
	}

	for i, tc := range testcases {
		actual := checkOverlap(tc.radius, tc.x, tc.y, tc.x1, tc.y1, tc.x2, tc.y2)
		if actual != tc.expected {
			t.Errorf("Testcase checkOverlap#%02d failed: want %t, got %t",
				i+1, tc.expected, actual)
		}
	}
}
