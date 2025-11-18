package lc0718

import "testing"

func Test_findLength(t *testing.T) {
	testcases := []struct {
		nums1    []int
		nums2    []int
		expected int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	}

	for i, tc := range testcases {
		actual := findLength(tc.nums1, tc.nums2)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.nums1, tc.nums2, tc.expected, actual)
		}
	}
}
