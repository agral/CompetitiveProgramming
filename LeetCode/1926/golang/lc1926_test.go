package lc1926

import "testing"

func Test_nearestExit(t *testing.T) {
	testcases := []struct {
		maze     [][]byte
		entrance []int
		expected int
	}{
		{[][]byte{
			{'+', '+', '.', '+'},
			{'.', '.', '.', '+'},
			{'+', '+', '+', '.'},
		}, []int{1, 2}, 1}, // need one step to reach [0, 2], which is the nearest exit.
		{[][]byte{
			{'+', '+', '+'},
			{'.', '.', '.'},
			{'+', '+', '+'},
		}, []int{1, 0}, 2}, // need two steps to reach [1, 2], which is the only exit.
		{[][]byte{
			{'.', '+'},
		}, []int{0, 0}, -1}, // there are no other exists in this maze.
	}

	for i, tc := range testcases {
		actual := nearestExit(tc.maze, tc.entrance)
		if actual != tc.expected {
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
