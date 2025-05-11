package lc0841

import "testing"

func Test_canVisitAllRooms(t *testing.T) {
	testcases := []struct {
		rooms    [][]int
		expected bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false}, // unable to visit room #2
	}

	for i, tc := range testcases {
		actual := canVisitAllRooms(tc.rooms)
		if actual != tc.expected {
			t.Errorf("Testcase %02d failed: want %t, got %t", i+1, tc.expected, actual)
		}
	}
}
