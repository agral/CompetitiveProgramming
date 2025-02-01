package lc0328

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	headEven := &ListNode{Val: -1, Next: nil}
	headOdd := &ListNode{Val: -1, Next: nil}

	currEven := headEven
	currOdd := headOdd
	isEven := false

	for head != nil {
		if isEven {
			currEven.Next = head
			currEven = head
		} else {
			currOdd.Next = head
			currOdd = head
		}

		head = head.Next
		isEven = !isEven
	}

	// Reconstruct the list: first all the odd nodes, then all the even nodes.
	currOdd.Next = headEven.Next
	currEven.Next = nil
	return headOdd.Next
}
