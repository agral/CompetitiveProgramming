package lc0611

import "testing"

func Test_canFormTriangle(t *testing.T) {
	testcases := []struct {
		first    int
		second   int
		third    int
		expected bool
	}{
		{1, 2, 2, true},
		{1, 2, 3, false},
		{1, 3, 4, false},
		{1, 4, 4, true},
	}

	for i, tc := range testcases {
		actual := canFormTriangle(tc.first, tc.second, tc.third)
		if actual != tc.expected {
			t.Errorf("Testcase canFormTriangle#%02d (%d, %d, %d) failed: want %t, got %t",
				i+1, tc.first, tc.second, tc.third, tc.expected, actual)
		}
	}
}

func Test_triangleNumber(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 2, 3, 4}, 3},
		{[]int{4, 2, 3, 4}, 4},
	}

	for i, tc := range testcases {
		actual := triangleNumber(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase triangleNumber#%02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
