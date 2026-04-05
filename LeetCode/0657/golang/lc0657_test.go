package lc0657

import "testing"

func Test_judgeCircle(t *testing.T) {
	testcases := []struct {
		moves    string
		expected bool
	}{
		{"UD", true},
		{"LL", false},
	}

	for i, tc := range testcases {
		actual := judgeCircle(tc.moves)
		if actual != tc.expected {
			t.Errorf("Testcase judgeCircle#%02d (%s) failed: want %t, got %t",
				i+1, tc.moves, tc.expected, actual)
		}
	}
}
