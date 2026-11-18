package lc2914

import "testing"

func Test_minChanges(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"1001", 2}, // e.g. "0011"
		{"10", 1},   // e.g. "11"
		{"0000", 0}, // already "beautiful"
	}

	for i, tc := range testcases {
		actual := minChanges(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase minChanges#%02d (%s) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
