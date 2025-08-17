package lc0837

import (
	"math"
	"testing"
)

func Test_new21Game(t *testing.T) {
	testcases := []struct {
		n        int
		k        int
		maxPts   int
		expected float64
	}{
		{10, 1, 10, 1.00000},
		{6, 1, 10, 0.60000},
		{21, 17, 10, 0.73278},
	}

	for i, tc := range testcases {
		actual := new21Game(tc.n, tc.k, tc.maxPts)
		if math.Abs(actual-tc.expected) > 0.00001 {
			t.Errorf("Testcase %02d (n=%d | k=%d | maxPts=%d) failed: want %f, got %f",
				i+1, tc.n, tc.k, tc.maxPts, tc.expected, actual)
		}
	}
}
