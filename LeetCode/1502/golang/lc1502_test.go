package lc1502

import "testing"

func Test_canMakeArithmeticProgression(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected bool
	}{
		{[]int{3, 5, 1}, true},
		{[]int{1, 2, 4}, false},
	}

	for i, tc := range testcases {
		actual := canMakeArithmeticProgression(tc.arr)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.arr, tc.expected, actual)
		}
	}
}
