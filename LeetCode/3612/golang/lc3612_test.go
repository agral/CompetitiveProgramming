package lc3612

import "testing"

func Test_processStr(t *testing.T) {
	testcases := []struct {
		s        string
		expected string
	}{
		{"a#b%*", "ba"},
		{"z*#", ""},
	}

	for i, tc := range testcases {
		actual := processStr(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase processStr#%02d (%s) failed: want %s, got %s",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
