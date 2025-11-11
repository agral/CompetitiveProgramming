package lc0474

import "testing"

func Test_findMaxForm(t *testing.T) {
	testcases := []struct {
		strs      []string
		numZeroes int
		numOnes   int
		expected  int
	}{
		{[]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		{[]string{"10", "0", "1"}, 1, 1, 2},
	}

	for i, tc := range testcases {
		actual := findMaxForm(tc.strs, tc.numZeroes, tc.numOnes)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | numZeroes=%d | numOnes=%d) failed: want %d, got %d",
				i+1, tc.strs, tc.numZeroes, tc.numOnes, tc.expected, actual)
		}
	}
}
