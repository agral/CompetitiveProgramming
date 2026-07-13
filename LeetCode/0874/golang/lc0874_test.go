package lc0874

import "testing"

func Test_robotSim(t *testing.T) {
	testcases := []struct {
		commands  []int
		obstacles [][]int
		expected  int
	}{
		{
			[]int{4, -1, 3},
			[][]int{},
			25,
		},
		{
			[]int{4, -1, 4, -2, 4},
			[][]int{{2, 4}},
			65,
		},
		{
			[]int{6, -1, -1, 6},
			[][]int{{0, 0}},
			36,
		},
	}

	for i, tc := range testcases {
		actual := robotSim(tc.commands, tc.obstacles)
		if actual != tc.expected {
			t.Errorf("Testcase robotSim#%02d (commands=%v, obstacles=%v) failed: want %d, got %d",
				i+1, tc.commands, tc.obstacles, tc.expected, actual)
		}
	}
}
