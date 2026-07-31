package lc3016

import "testing"

func Test_minimumPushes(t *testing.T) {
	testcases := []struct {
		word     string
		expected int
	}{
		{"abcde", 5},
		{"xyzxyzxyzxyz", 12},
		{"aabbccddeeffgghhiiiiii", 24},
	}

	for i, tc := range testcases {
		actual := minimumPushes(tc.word)
		if actual != tc.expected {
			t.Errorf("Testcase minimumPushes#%02d (%v) failed: want %d, got %d",
				i+1, tc.word, tc.expected, actual)
		}
	}
}
