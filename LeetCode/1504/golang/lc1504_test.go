package lc1504

import "testing"

func Test_numSubmat(t *testing.T) {
	testcases := []struct {
		mat      [][]int
		expected int
	}{
		{[][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}, 13},
		{[][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}, 24},
	}

	for i, tc := range testcases {
		actual := numSubmat(tc.mat)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.mat, tc.expected, actual)
		}
	}
}
