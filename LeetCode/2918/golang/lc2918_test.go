package lc2918

import "testing"

func Test_minSum(t *testing.T) {
	testcases := []struct {
		nums1    []int
		nums2    []int
		expected int64
	}{
		{[]int{3, 2, 0, 1, 0}, []int{6, 5, 0}, 12},
		{[]int{2, 0, 2, 0}, []int{1, 4}, -1},
	}

	for i, tc := range testcases {
		actual := minSum(tc.nums1, tc.nums2)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v|%v) failed: want %d, got %d",
				i+1, tc.nums1, tc.nums2, tc.expected, actual)
		}
	}
}
