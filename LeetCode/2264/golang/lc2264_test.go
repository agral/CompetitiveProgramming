package lc2264

import "testing"

func Test_largestGoodInteger(t *testing.T) {
	testcases := []struct {
		num      string
		expected string
	}{
		{"6777133339", "777"},
		{"2300019", "000"},
		{"42352338", ""},
	}

	for i, tc := range testcases {
		actual := largestGoodInteger(tc.num)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %s, got %s", i+1, tc.num, tc.expected, actual)
		}
	}
}
