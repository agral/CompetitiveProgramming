package lc0371

import "testing"

func Test_getSum(t *testing.T) {
	testcases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{7, 13, 20},
		{-7, 13, 6},
	}

	for i, tc := range testcases {
		actual := getSum(tc.a, tc.b)
		if actual != tc.expected {
			t.Errorf("Testcase getSum#%02d (%v) failed: want %d, got %d",
				i+1, tc.a, tc.expected, actual)
		}
	}
}
