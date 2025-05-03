package lc1006

import "testing"

func Test_clumsy(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{1, 1}, // 1 == 1
		{2, 2}, // 2 * 1 == 2
		{3, 6}, // 3 * 2 / 1 == 6
		{4, 7}, // 4 * 3 / 2 + 1 == 7
		{5, 7}, // 5 * 4 / 3 + 2 - 1 == 7
		{6, 8}, // 6 * 5 / 4 + 3 - 2 * 1 == 8
		{7, 6}, // 7 * 6 / 5 + 4 - 3 * 2 / 1 == 6
		{8, 9}, // 8 * 7 / 6 + 5 - 4 * 3 / 2 + 1 == 9
		{9, 11},
		{10, 12},
		{11, 10},
		{12, 13},
	}

	for i, tc := range testcases {
		actual := clumsy(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
