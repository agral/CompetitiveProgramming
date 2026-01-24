package lc2485

import "testing"

func Test_sumNums(t *testing.T) {
	testcases := []struct {
		start    int
		stop     int
		expected int
	}{
		{1, 6, 21},
		{6, 8, 21},
		{1, 1, 1},
		{4, 4, 4},
		{1, 100, 5050},
	}
	for i, tc := range testcases {
		actual := sumNums(tc.start, tc.stop)
		if actual != tc.expected {
			t.Errorf("Testcase sumNums#%02d (%d-%d) failed: want %d, got %d",
				i+1, tc.start, tc.stop, tc.expected, actual)
		}
	}
}

func Test_pivotInteger(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, -1},
		{3, -1},
		{4, -1},
		{5, -1},
		{6, -1},
		{7, -1},
		{8, 6},
		{9, -1},
		{49, 35},
		{288, 204},
		{289, -1},
	}

	for i, tc := range testcases {
		actual := pivotInteger(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase pivotInteger#%02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
