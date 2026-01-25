package lc1984

import "testing"

func Test_minimumDifference(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{90}, 1, 0},
		{[]int{9, 4, 1, 7}, 2, 2},
	}

	for i, tc := range testcases {
		actual := minimumDifference(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase minimumDifference#%02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
