package lc0944

import "testing"

func Test_minDeletionSize(t *testing.T) {
	testcases := []struct {
		strs     []string
		expected int
	}{
		{[]string{"abc", "bce", "cae"}, 1},
		{[]string{"a", "b"}, 0},
		{[]string{"zyx", "wvu", "tsr"}, 3},
	}

	for i, tc := range testcases {
		actual := minDeletionSize(tc.strs)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.strs, tc.expected, actual)
		}
	}
}
