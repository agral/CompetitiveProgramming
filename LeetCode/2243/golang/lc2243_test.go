package lc2243

import "testing"

func Test_digitSum(t *testing.T) {
	testcases := []struct {
		s        string
		k        int
		expected string
	}{
		{"11111222223", 3, "135"},
		{"00000000", 3, "000"},
	}

	for i, tc := range testcases {
		actual := digitSum(tc.s, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase digitSum#%02d (s=%v, k=%d) failed: want %s, got %s",
				i+1, tc.s, tc.k, tc.expected, actual)
		}
	}
}
