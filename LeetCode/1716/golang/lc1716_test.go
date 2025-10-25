package lc1716

import "testing"

func Test_totalMoney(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{4, 10},
		{10, 37},
		{20, 96},
	}

	for i, tc := range testcases {
		actual := totalMoney(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
