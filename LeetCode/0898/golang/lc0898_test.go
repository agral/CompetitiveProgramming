package lc0898

import "testing"

func Test_subarrayBitwiseORs(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected int
	}{
		{[]int{0}, 1},
		{[]int{1, 1, 2}, 3},
		{[]int{1, 2, 4}, 6},
	}

	for i, tc := range testcases {
		actual := subarrayBitwiseORs(tc.arr)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.arr, tc.expected, actual)
		}
	}
}
