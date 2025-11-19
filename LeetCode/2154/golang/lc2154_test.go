package lc2154

import "testing"

func Test_findFinalValue(t *testing.T) {
	testcases := []struct {
		input    []int
		original int
		expected int
	}{
		{[]int{5, 3, 6, 1, 12}, 3, 24},
		{[]int{2, 7, 9}, 4, 4},
	}

	for i, tc := range testcases {
		actual := findFinalValue(tc.input, tc.original)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.input, tc.original, tc.expected, actual)
		}
	}
}
