package lc2125

import "testing"

func Test_numberOfBeams(t *testing.T) {
	testcases := []struct {
		bank     []string
		expected int
	}{
		{[]string{"011001", "000000", "010100", "001000"}, 8},
		{[]string{"000", "111", "000"}, 0},
	}

	for i, tc := range testcases {
		actual := numberOfBeams(tc.bank)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.bank, tc.expected, actual)
		}
	}
}
