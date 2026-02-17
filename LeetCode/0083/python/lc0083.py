from typing import List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: easy+
# Solved on: 2026-02-16
class Solution:
    def deleteDuplicates(self, head: ListNode | None) -> ListNode | None:
        if not head or not head.next:
            return head
        prev = head
        curr = head.next
        while prev is not None and curr is not None:
            while curr is not None and prev.val == curr.val:
                # Remove current node from the linked list - it is a duplicate of prev.
                prev.next = curr.next
                curr = curr.next
            prev = prev.next
            curr = prev.next if prev is not None else None
        return head
