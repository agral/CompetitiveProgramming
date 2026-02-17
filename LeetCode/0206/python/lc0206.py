class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy
# Solved on: 2026-02-17
class Solution:
    def reverseList(self, head: ListNode | None) -> ListNode | None:
        # Don't reverse the list if it has less than two nodes:
        if head is None or head.next is None:
            return head

        # Reverse the "inside" of the list by keeping three pointers: prev, curr & next:
        prev = head
        curr = head.next
        next = curr.next
        while next is not None:
            curr.next = prev
            prev = curr
            curr = next
            next = curr.next

        # Reverse the first and last nodes, set the former tail of the list as a head and return it:
        head.next = None
        curr.next = prev
        head = curr
        return head
