package lc2425

import "testing"

func TestXorAllNums(t *testing.T) {
	testcases := []struct {
		nums1    []int
		nums2    []int
		expected int
	}{
		{[]int{2, 1, 3}, []int{10, 2, 5, 0}, 13},
		{[]int{1, 2}, []int{3, 4}, 0},
		{[]int{9, 8, 7, 6}, []int{1, 2, 3, 4}, 0},
		{[]int{9, 8, 7, 6}, []int{1, 2, 3, 4, 5, 6}, 0},
		{[]int{9, 8, 7}, []int{1, 2, 3}, 6},
		{[]int{9, 8, 7}, []int{1, 2, 3, 33, 333}, 362},
		{[]int{1, 11, 111, 1111}, []int{1, 2, 3, 33, 333}, 1074},
		{[]int{1, 11, 111, 1111, 11111}, []int{3, 33, 333, 3333}, 3178},
	}

	for i, tc := range testcases {
		actual := xorAllNums(tc.nums1, tc.nums2)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.nums1, tc.nums2, tc.expected, actual)
		}
	}
}
