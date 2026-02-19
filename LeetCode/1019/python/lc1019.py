from typing import List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Runtime complexity: O(n)
# Auxiliary space complexity: O(n)
# Subjective level: medium
# Solved on: 2026-02-19
class Solution:
    def nextLargerNodes(self, head: ListNode | None) -> List[int]:
        stack = []
        values = []
        curr = head
        while curr != None:
            while stack and curr.val > values[stack[-1]]:
                idx = stack.pop()
                values[idx] = curr.val
            stack.append(len(values))
            values.append(curr.val)
            curr = curr.next

        for local_maximum_idx in stack:
            values[local_maximum_idx] = 0

        return values
