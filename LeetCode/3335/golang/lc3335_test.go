package lc3335

import "testing"

func Test_lengthAfterTransformations(t *testing.T) {
	testcases := []struct {
		s        string
		t        int
		expected int
	}{
		{"abcyy", 2, 7}, // "cdeabab"
		{"azbk", 1, 5},  // "babcl"
	}

	for i, tc := range testcases {
		actual := lengthAfterTransformations(tc.s, tc.t)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %d, got %d", i+1, tc.s, tc.expected, actual)
		}
	}
}
