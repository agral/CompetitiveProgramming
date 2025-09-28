package lc0976

import "testing"

func Test_largestPerimeter(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 1, 2}, 5},
		{[]int{1, 1, 2, 10}, 0},
	}

	for i, tc := range testcases {
		actual := largestPerimeter(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
