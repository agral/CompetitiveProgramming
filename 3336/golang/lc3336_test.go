package lc3336

import "testing"

func Test_subsequencePairCount(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{10, 20, 30}, 2},
		{[]int{1, 1, 1, 1}, 50},
	}

	for i, tc := range testcases {
		actual := subsequencePairCount(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase subsequencePairCount#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
