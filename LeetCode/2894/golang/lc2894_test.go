package lc2894

import "testing"

func Test_differenceOfSums(t *testing.T) {
	testcases := []struct {
		n        int
		m        int
		expected int
	}{
		{10, 3, 19}, // 37 - 18 == 19
		{5, 6, 15},  // 15 - 0 == 15
		{5, 1, -15}, // 0 - 15 == -15
	}

	for i, tc := range testcases {
		actual := differenceOfSums(tc.n, tc.m)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d | %d) failed: want %d, got %d",
				i+1, tc.n, tc.m, tc.expected, actual)
		}
	}
}
