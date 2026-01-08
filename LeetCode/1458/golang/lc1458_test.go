package lc1458

import "testing"

func Test_maxDotProduct(t *testing.T) {
	testcases := []struct {
		nums1    []int
		nums2    []int
		expected int
	}{
		{[]int{2, 1, -2, 5}, []int{3, 0, -6}, 18},
		{[]int{3, -2}, []int{2, -6, 7}, 21},
		{[]int{-1, -1}, []int{1, 1}, -1},
	}

	for i, tc := range testcases {
		actual := maxDotProduct(tc.nums1, tc.nums2)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.nums1, tc.nums2, tc.expected, actual)
		}
	}
}
