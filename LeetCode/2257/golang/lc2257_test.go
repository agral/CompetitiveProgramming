package lc2257

import "testing"

func Test_countUnguarded(t *testing.T) {
	testcases := []struct {
		H        int
		W        int
		guards   [][]int
		walls    [][]int
		expected int
	}{
		{
			/*
			 * GW.+..
			 * +G++W.
			 * ++WG++
			 * ++.+..
			 */
			/*H=*/ 4,
			/*W=*/ 6,
			/*guards=*/ [][]int{{0, 0}, {1, 1}, {2, 3}},
			/*walls=*/ [][]int{{0, 1}, {2, 2}, {1, 4}},
			/*expected=*/ 7,
		},
		{
			/*
			 * .W.
			 * WGW
			 * .W.
			 */
			/*H=*/ 3,
			/*W=*/ 3,
			/*guards=*/ [][]int{{1, 1}},
			/*walls=*/ [][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}},
			/*expected=*/ 4,
		},
	}

	for i, tc := range testcases {
		actual := countUnguarded(tc.H, tc.W, tc.guards, tc.walls)
		if actual != tc.expected {
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
