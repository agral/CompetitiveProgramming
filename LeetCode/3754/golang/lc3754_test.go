package lc3754

import "testing"

func Test_sumAndMultiply(t *testing.T) {
	testcases := []struct {
		n        int
		expected int64
	}{
		{10203004, int64(12340)},
		{1000, int64(1)},
		{0, int64(0)},
	}

	for i, tc := range testcases {
		actual := sumAndMultiply(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase sumAndMultiply#%02d (%v) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
