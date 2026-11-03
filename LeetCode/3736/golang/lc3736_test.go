package lc3736

import "testing"

func Test_minMoves(t *testing.T) {
	testcases := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 1, 3}, 3},
		{[]int{4, 4, 5}, 2},
	}

	for i, tc := range testcases {
		actual := minMoves(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase minMoves#%02d (%v) failed: want %d, got %d",
				i+1, tc.input, tc.expected, actual)
		}
	}
}
