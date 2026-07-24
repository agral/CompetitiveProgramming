package lc3514

import "testing"

func Test_uniqueXorTriplets(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3}, 2},
		{[]int{6, 7, 8, 9}, 4},
	}

	for i, tc := range testcases {
		actual := uniqueXorTriplets(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase uniqueXorTriplets#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
