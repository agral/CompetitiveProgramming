package lc0206

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

func reverseList(head *ListNode) *ListNode {
	// Don't reverse lists containing 0 or 1 nodes:
	if head == nil || head.Next == nil {
		return head
	}

	// Reverse the main body of the list by keeping three pointers; curr, prev, and next:
	prev := head
	curr := head.Next
	next := curr.Next
	for next != nil {
		curr.Next = prev
		prev = curr
		curr = next
		next = curr.Next
	}

	// Finally reverse the last link and update the list's head:
	head.Next = nil // the old list's head is the reversed list's tail; its Next element should be nil.
	curr.Next = prev
	head = curr
	return head
}

/* Don't really want to write unit tests for this. It works.
func main() {
	myList := makeList([]int{1, 2, 3, 4, 5})
	printList(myList)
	reversed := reverseList(myList)
	printList(reversed)
}
*/
