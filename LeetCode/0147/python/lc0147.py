class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Runtime complexity: O(n^2)
# Auxiliary space complexity: O(1) (reusing the existing nodes)
# Subjective level: medium
# Solved on: 2026-03-09
class Solution:
    def insertionSortList(self, head: ListNode | None) -> ListNode | None:
        ans = ListNode(-42000) # First node will not be part of the answer;
        # the value for this "extra" node is chosen so that it will be lower
        # than any value in the list -> and so it will be "sorted" first.

        prev = ans # prev holds the tail of the sorted list

        curr = head # curr holds the currently processed node in the original list.
        while curr:
            next = curr.next
            if prev.val >= curr.val:
                prev = ans
            while prev.next and prev.next.val < curr.val: # find the spot in the sorted list (ans)
                prev = prev.next   # to put this node into.
            curr.next = prev.next
            prev.next = curr

            curr = next

        return ans.next
