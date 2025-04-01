package lc2140

import "testing"

func Test_mostPoints(t *testing.T) {
	testcases := []struct {
		questions [][]int
		expected  int64
	}{
		{[][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}, 5},
		{[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, 7},
	}

	for i, tc := range testcases {
		actual := mostPoints(tc.questions)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.questions, tc.expected, actual)
		}
	}
}
