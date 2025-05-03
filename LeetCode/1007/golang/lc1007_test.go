package lc1007

import "testing"

func Test_minDominoRotations(t *testing.T) {
	testcases := []struct {
		tops     []int
		bottoms  []int
		expected int
	}{
		{[]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}, 2},
		{[]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}, -1},
	}

	for i, tc := range testcases {
		actual := minDominoRotations(tc.tops, tc.bottoms)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v|%v) failed: want %d, got %d",
				i+1, tc.tops, tc.bottoms, tc.expected, actual)
		}
	}
}
