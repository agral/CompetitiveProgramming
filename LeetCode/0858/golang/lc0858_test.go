package lc0858

import "testing"

func Test_mirrorReflection(t *testing.T) {
	testcases := []struct {
		p        int
		q        int
		expected int
	}{
		{2, 1, 2},
		{3, 1, 1},
		{3, 2, 0},
	}

	for i, tc := range testcases {
		actual := mirrorReflection(tc.p, tc.q)
		if actual != tc.expected {
			t.Errorf("Testcase mirrorReflection#%02d (p=%d, q=%d) failed: want %d, got %d",
				i+1, tc.p, tc.q, tc.expected, actual)
		}
	}
}
