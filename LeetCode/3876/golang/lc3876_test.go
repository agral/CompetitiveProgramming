package lc3876

import "testing"

func Test_uniformArray(t *testing.T) {
	testcases := []struct {
		nums1    []int
		expected bool
	}{
		{[]int{1, 4, 7}, true}, // [1, 4-1=3, 7]
		{[]int{2, 3}, false},   // 2-3== -1, can't make the array (-1 < 1).
		{[]int{4, 6}, true},    // already of same parity, just rewrite nums1 to nums2.
	}

	for i, tc := range testcases {
		actual := uniformArray(tc.nums1)
		if actual != tc.expected {
			t.Errorf("Testcase uniformArray#%02d (%v) failed: want %t, got %t",
				i+1, tc.nums1, tc.expected, actual)
		}
	}
}
