package lc3300

import "testing"

func Test_minElement(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 12, 13, 14}, 1},
		{[]int{1, 2, 3, 4}, 1},
		{[]int{999, 19, 199}, 10},
	}

	for i, tc := range testcases {
		actual := minElement(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase minElement#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
