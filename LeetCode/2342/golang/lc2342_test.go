package lc2342

import "testing"

func Test_getDigitSum(t *testing.T) {
	testcases := []struct {
		num      int
		expected int
	}{
		{1, 1},
		{9, 9},
		{13, 4},
		{101, 2},
		{12345, 15},
		{999999999, 81},
	}

	for i, tc := range testcases {
		actual := getDigitSum(tc.num)
		if actual != tc.expected {
			t.Errorf("Testcase getDigitSum#%02d (%d) failed: want %d, got %d", i+1, tc.num, tc.expected, actual)
		}
	}
}

func Test_maximumSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{18, 43, 36, 13, 7}, 54},
		{[]int{10, 12, 19, 14}, -1},
	}

	for i, tc := range testcases {
		actual := maximumSum(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
