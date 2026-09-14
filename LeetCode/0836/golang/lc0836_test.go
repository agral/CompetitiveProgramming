package lc0836

import "testing"

func Test_isRectangleOverlap(t *testing.T) {
	testcases := []struct {
		rec1     []int
		rec2     []int
		expected bool
	}{
		{[]int{0, 0, 2, 2}, []int{1, 1, 3, 3}, true},
		{[]int{0, 0, 1, 1}, []int{1, 0, 2, 1}, false},
		{[]int{0, 0, 1, 1}, []int{2, 2, 3, 3}, false},
	}

	for i, tc := range testcases {
		actual := isRectangleOverlap(tc.rec1, tc.rec2)
		if actual != tc.expected {
			t.Errorf("Testcase isRectangleOverlap#%02d (rec1=%v, rec2=%v) failed: want %t, got %t",
				i+1, tc.rec1, tc.rec2, tc.expected, actual)
		}
	}
}
