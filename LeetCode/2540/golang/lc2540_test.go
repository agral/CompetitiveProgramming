package lc2540

import "testing"

func Test_getCommon(t *testing.T) {
	testcases := []struct {
		nums1    []int
		nums2    []int
		expected int
	}{
		{[]int{1, 2, 3}, []int{2, 4}, 2},           // the smallest (and only) common element to both slices is 2.
		{[]int{1, 2, 3, 6}, []int{2, 3, 4, 5}, 2},  // 2 and 3 match; 2 is the smallest.
		{[]int{1, 3, 5, 7}, []int{2, 4, 6, 8}, -1}, // no matches.
	}

	for i, tc := range testcases {
		actual := getCommon(tc.nums1, tc.nums2)
		if actual != tc.expected {
			t.Errorf("Testcase getCommon#%02d (%v | %v) failed: want %d, got %d",
				i+1, tc.nums1, tc.nums2, tc.expected, actual)
		}
	}
}
