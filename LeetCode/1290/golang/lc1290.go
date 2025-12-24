package lc1290

type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2025-12-24
func getDecimalValue(head *ListNode) int {
	ans := 0
	for head != nil {
		ans *= 2
		ans += head.Val
		head = head.Next
	}
	return ans
}
