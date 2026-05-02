package lc0788

import "testing"

func Test_isGood(t *testing.T) {
	testcases := []struct {
		i        int
		expected bool
	}{
		{1, false},
		{12, true},
		{2, true},
		{3, false},
		{4, false},
		{5, true},
		{50, true},
		{500, true},
		{503, false},
		{508, true},
	}
	for i, tc := range testcases {
		actual := isGood(tc.i)
		if actual != tc.expected {
			t.Errorf("Testcase rotatedDigits#%02d (%d) failed: want %t, got %t",
				i+1, tc.i, tc.expected, actual)
		}
	}
}

func Test_rotatedDigits(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{10, 4},
		{1, 0},
		{2, 1},
	}

	for i, tc := range testcases {
		actual := rotatedDigits(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase rotatedDigits#%02d (%d) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
