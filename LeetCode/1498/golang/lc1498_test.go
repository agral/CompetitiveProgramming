package lc1498

import "testing"

func Test_numSubseq(t *testing.T) {
	testcases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{3, 5, 6, 7}, 9, 4},
		{[]int{3, 3, 6, 8}, 10, 6},
		{[]int{2, 3, 3, 4, 6, 7}, 12, 61},
	}

	for i, tc := range testcases {
		actual := numSubseq(tc.nums, tc.target)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.target, tc.expected, actual)
		}
	}
}
