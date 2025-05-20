package lc3355

import "testing"

func Test_isZeroArray(t *testing.T) {
	testcases := []struct {
		nums     []int
		queries  [][]int
		expected bool
	}{
		{[]int{1, 0, 1}, [][]int{{0, 2}}, true},
		{[]int{4, 3, 2, 1}, [][]int{{1, 3}, {0, 2}}, false},
	}

	for i, tc := range testcases {
		actual := isZeroArray(tc.nums, tc.queries)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %t, got %t",
				i+1, tc.nums, tc.queries, tc.expected, actual)
		}
	}
}
