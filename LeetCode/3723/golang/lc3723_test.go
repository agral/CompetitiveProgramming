package lc3723

import "testing"

func Test_maxSumOfSquares(t *testing.T) {
	testcases := []struct {
		num      int
		sum      int
		expected string
	}{
		{2, 3, "30"},
		{2, 17, "98"},
		{1, 10, ""},
	}

	for i, tc := range testcases {
		actual := maxSumOfSquares(tc.num, tc.sum)
		if actual != tc.expected {
			t.Errorf("Testcase maxSumOfSquares#%02d (num=%d, sum=%d) failed: want %s, got %s",
				i+1, tc.num, tc.sum, tc.expected, actual)
		}
	}
}
