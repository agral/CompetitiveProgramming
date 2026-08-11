package lc2996

import "testing"

func Test_missingInteger(t *testing.T) {
	testcases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 2, 5}, 6},
		{[]int{3, 4, 5, 1, 12, 14, 13}, 15},
		{[]int{29, 30, 31, 32, 33, 34, 35, 36, 37}, 297},
		{[]int{4, 5, 6, 7, 8, 8, 9, 4, 3, 2, 7}, 30},      // WA #01, forgot to reset the CSP after no-seq.
		{[]int{14, 9, 6, 9, 7, 9, 10, 4, 9, 9, 4, 4}, 15}, // What! How?! It should have been 19...
		// Ah, okay. So the prefix sequence somehow has to include nums[0]...? Maybe? Why?
	}

	for i, tc := range testcases {
		actual := missingInteger(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase missingInteger#%02d (%v) failed: want %d, got %d",
				i+1, tc.input, tc.expected, actual)
		}
	}
}
