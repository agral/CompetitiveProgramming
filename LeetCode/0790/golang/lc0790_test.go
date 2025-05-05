package lc0790

import "testing"

func Test_numTilings(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 5},
		{4, 11},
		{5, 24},
		{6, 53},
		{7, 117},
		{1000, 979232805},
	}

	for i, tc := range testcases {
		actual := numTilings(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
