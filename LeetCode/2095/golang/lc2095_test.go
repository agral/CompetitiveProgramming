package lc2095

import "testing"

func isEqual(lhs *ListNode, rhs *ListNode) bool {
	for lhs != nil && rhs != nil {
		if lhs.Val != rhs.Val {
			return false
		}
		lhs = lhs.Next
		rhs = rhs.Next
	}
	return lhs == nil && rhs == nil
}

func Test_deleteMiddle(t *testing.T) {
	testcases := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{makeList([]int{1, 3, 4, 7, 1, 2, 6}), makeList([]int{1, 3, 4, 1, 2, 6})},
		{makeList([]int{1, 2, 3, 4}), makeList([]int{1, 2, 4})},
		{makeList([]int{2, 1}), makeList([]int{2})},
		{makeList([]int{1}), makeList([]int{})},
	}

	for i, tc := range testcases {
		actual := deleteMiddle(tc.input)
		if !isEqual(actual, tc.expected) {
			t.Errorf("Testcase %02d failed", i+1)
			printList(tc.input)
		}
	}
}
