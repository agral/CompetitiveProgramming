package lc0328

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0], Next: nil}
	prev := head
	for i := 1; i < len(vals); i++ {
		newNode := &ListNode{Val: vals[i], Next: nil}
		prev.Next = newNode
		prev = newNode
	}
	return head
}

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println("nil")
}

func IsEqual(lhs *ListNode, rhs *ListNode) bool {
	for lhs != nil && rhs != nil {
		if lhs.Val != rhs.Val {
			return false
		}
		lhs = lhs.Next
		rhs = rhs.Next
	}
	return lhs == nil && rhs == nil
}
