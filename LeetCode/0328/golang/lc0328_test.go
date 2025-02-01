package lc0328

import "testing"

func Test_oddEvenList(t *testing.T) {
	testcases := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{makeList([]int{1, 2, 3, 4, 5}), makeList([]int{1, 3, 5, 2, 4})},
		{makeList([]int{2, 1, 3, 5, 6, 4, 7}), makeList([]int{2, 3, 6, 7, 1, 5, 4})},
	}

	for i, tc := range testcases {
		actual := oddEvenList(tc.input)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase %02d failed. Actual and expected contents follow.", i+1)
			printList(actual)
			printList(tc.expected)
		}
	}
}
