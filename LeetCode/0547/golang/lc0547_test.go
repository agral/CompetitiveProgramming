package lc0547

import "testing"

func Test_findCircleNum(t *testing.T) {
	testcases := []struct {
		isConnected [][]int
		expected    int
	}{
		{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
		{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
	}

	for i, tc := range testcases {
		actual := findCircleNum(tc.isConnected)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.isConnected, tc.expected, actual)
		}
	}
}
