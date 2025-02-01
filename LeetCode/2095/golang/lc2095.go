package lc2095

func deleteMiddle(head *ListNode) *ListNode {
	// Iterate over the entire list to calculate its length:
	length := 0
	curr := head
	for curr != nil {
		curr = curr.Next
		length++
	}

	if length <= 1 {
		return nil
	}

	// Then move over to just before the index to be deleted.
	// Note: this might be done in one batch with calculating the length
	// where the "mid" index would be moving on every two moves of "length" index;
	// but it's less readable and results in the same complexity as this solution.
	// So this solution calculates length first; then locates the middle index.
	rm := length / 2
	curr = head
	for i := 0; i < rm-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next

	return head
}
