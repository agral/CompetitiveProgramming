package lc1545

import "testing"

func Test_findKthBit(t *testing.T) {
	testcases := []struct {
		n        int
		k        int
		expected byte
	}{
		{3, 1, '0'},  // [0]111001
		{4, 11, '1'}, // 0111001101[1]0001
	}

	for i, tc := range testcases {
		actual := findKthBit(tc.n, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase findKthBit#%02d (%d, %d) failed: want %c, got %c",
				i+1, tc.n, tc.k, tc.expected, actual)
		}
	}
}
