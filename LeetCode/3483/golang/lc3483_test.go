package lc3483

import "testing"

func Test_totalNumbers(t *testing.T) {
	testcases := []struct {
		digits   []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 12},
		{[]int{0, 2, 2}, 2}, // just 202 and 220. 022 is invalid.
		{[]int{6, 6, 6}, 1},
		{[]int{1, 3, 5}, 0}, // no even numbers can be formed.
	}

	for i, tc := range testcases {
		actual := totalNumbers(tc.digits)
		if actual != tc.expected {
			t.Errorf("Testcase totalNumbers#%02d (%v) failed: want %d, got %d",
				i+1, tc.digits, tc.expected, actual)
		}
	}
}
