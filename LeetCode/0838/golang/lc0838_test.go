package lc0838

import "testing"

func Test_pushDominoes(t *testing.T) {
	testcases := []struct {
		dominoes string
		expected string
	}{
		{"RR.L", "RR.L"},
		{".L.R...LR..L..", "LL.RR.LLRRLL.."},
	}

	for i, tc := range testcases {
		actual := pushDominoes(tc.dominoes)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %s, got %s", i+1, tc.dominoes, tc.expected, actual)
		}
	}
}
