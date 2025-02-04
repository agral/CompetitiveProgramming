package lc2130

func pairSum(head *ListNode) int {
	slow := head
	fast := head

	// iterate over the list, so that `slow` points to the middle of the list:
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	ans := 0
	// `revHalf` points to the head of the reversed second half of the original list
	// (which now starts at `revHalf` and ends at `slow`)
	revHalf := Reverse(slow)
	for revHalf != nil {
		ans = max(ans, head.Val+revHalf.Val)
		head = head.Next
		revHalf = revHalf.Next
	}
	return ans
}
