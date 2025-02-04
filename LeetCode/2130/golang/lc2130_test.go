package lc2130

import "testing"

func Test_pairSum(t *testing.T) {
	testcases := []struct {
		head     *ListNode
		expected int
	}{
		{makeList([]int{5, 4, 2, 1}), 6},
		{makeList([]int{4, 2, 2, 3}), 7},
		{makeList([]int{1, 100000}), 100001},
	}

	for i, tc := range testcases {
		actual := pairSum(tc.head)
		if actual != tc.expected {
			printList(tc.head)
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
