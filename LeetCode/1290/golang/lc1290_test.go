package lc1290

import "testing"

func MakeLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := ListNode{
		Val:  values[0],
		Next: nil,
	}
	var prev *ListNode = &head
	for i := 1; i < len(values); i++ {
		curr := ListNode{
			Val:  values[i],
			Next: nil,
		}
		prev.Next = &curr
		prev = &curr
	}
	return &head
}

func ListToSlice(l *ListNode) []int {
	ans := []int{}
	for l != nil {
		ans = append(ans, l.Val)
		l = l.Next
	}
	return ans
}

func Test_getDecimalValue(t *testing.T) {
	testcases := []struct {
		head     *ListNode
		expected int
	}{
		{MakeLinkedList([]int{1, 0, 1}), 5},
		{MakeLinkedList([]int{0}), 0},
		{MakeLinkedList([]int{1, 0, 1, 0}), 10},
	}

	for i, tc := range testcases {
		actual := getDecimalValue(tc.head)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d",
				i+1, ListToSlice(tc.head), tc.expected, actual)
		}
	}
}
