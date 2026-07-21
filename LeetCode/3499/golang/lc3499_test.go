package lc3499

import "testing"

func Test_maxActiveSectionsAfterTrade(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"01", 1},
		{"0100", 4},
		{"1000100", 7},
		{"01010", 4},
	}

	for i, tc := range testcases {
		actual := maxActiveSectionsAfterTrade(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase maxActiveSectionsAfterTrade#%02d (%s) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
