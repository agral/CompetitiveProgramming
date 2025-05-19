package lc3024

import "testing"

func Test_triangleType(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected string
	}{
		{[]int{3, 3, 3}, "equilateral"},
		{[]int{3, 4, 5}, "scalene"},
		{[]int{1, 2, 3}, "none"},
		{[]int{7, 4, 7}, "isosceles"},
		{[]int{7, 101, 7}, "none"},
	}

	for i, tc := range testcases {
		actual := triangleType(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %s, got %s", i+1, tc.nums, tc.expected, actual)
		}
	}
}
