package lc1399

import "testing"

func Test_countLargestGroup(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{13, 4},
		{2, 2},
		{14, 5},
		{15, 6},
		{18, 9}, // there's 9 groups, each of these has 2 members.
		{19, 9}, // there's 9 groups from the above + the 10th group containing sole "19".
		{20, 1}, // there's one group with three members: [2, 11, 20].
	}

	for i, tc := range testcases {
		actual := countLargestGroup(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}

func Test_getDecimalDigitsSum(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{9, 9},
		{10, 1},
		{11, 2},
		{111, 3},
		{1000, 1},
	}

	for _, tc := range testcases {
		actual := getDecimalDigitsSum(tc.n)
		if actual != tc.expected {
			t.Errorf("getDecimalDigitsSum(%d) returned %d (want: %d)", tc.n, actual, tc.expected)
		}
	}
}
