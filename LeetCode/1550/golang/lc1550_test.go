package lc1550

import "testing"

func Test_threeConsecutiveOdds(t *testing.T) {
	testcases := []struct {
		input    []int
		expected bool
	}{
		{[]int{2, 6, 4, 1}, false},
		{[]int{1, 2, 34, 3, 4, 5, 7, 23, 12}, true}, // 5, 7, 23
	}

	for i, tc := range testcases {
		actual := threeConsecutiveOdds(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.input, tc.expected, actual)
		}
	}
}
