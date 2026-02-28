package lc1680

import "testing"

func Test_concatenatedBinary(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 6},
		{3, 27},
		{12, 505379714},
	}

	for i, tc := range testcases {
		actual := concatenatedBinary(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase concatenatedBinary#%02d (%d) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
