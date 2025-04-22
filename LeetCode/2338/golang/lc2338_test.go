package lc2338

import "testing"

func Test_idealArrays(t *testing.T) {
	testcases := []struct {
		n        int
		maxValue int
		expected int
	}{
		{2, 5, 10},
		{5, 3, 11},
	}

	for i, tc := range testcases {
		actual := idealArrays(tc.n, tc.maxValue)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (n=%d, maxValue=%d) failed: want %d, got %d",
				i+1, tc.n, tc.maxValue, tc.expected, actual)
		}
	}
}
