class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Runtime complexity: O(n)
# Auxiliary space complexity: O(1)
# Subjective level: medium-
# Solved on: 2026-02-17
class Solution:
    def oddEvenList(self, head: ListNode | None) -> ListNode | None:
        if head is None or head.next is None:
            return head

        headOdd, headEven = head, head.next
        currOdd, currEven = headOdd, headEven
        head = head.next.next
        isOdd = True # Linked list elements are 1-indexed, and we're currently at head+2 (index 3).
        while head is not None:
            if isOdd:
                currOdd.next = head
                currOdd = head
            else:
                currEven.next = head
                currEven = head
            head = head.next
            isOdd = not isOdd

        # Finally link the even nodes after all the odd nodes:
        currOdd.next = headEven
        currEven.next = None
        return headOdd
