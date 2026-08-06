package lc3345

import "testing"

func Test_smallestNumber(t *testing.T) {
	testcases := []struct {
		n        int
		t        int
		expected int
	}{
		{10, 2, 10},
		{15, 3, 16},
	}

	for i, tc := range testcases {
		actual := smallestNumber(tc.n, tc.t)
		if actual != tc.expected {
			t.Errorf("Testcase smallestNumber#%02d (n=%d, t=%d) failed: want %d, got %d",
				i+1, tc.n, tc.t, tc.expected, actual)
		}
	}
}

func Test_digitProduct(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{3, 3},
		{7, 7},
		{10, 0},
		{17, 7},
		{20, 0},
		{21, 2},
		{22, 4},
		{29, 18},
		{99, 81},
		{100, 0},
		{109, 0},
		{111, 1},
	}

	for i, tc := range testcases {
		actual := digitProduct(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase digitProduct #%02d (%d) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
