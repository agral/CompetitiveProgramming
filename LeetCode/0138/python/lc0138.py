class Node:
    def __init__(self, val=0, next=None, random=None):
        self.val = val
        self.next = next
        self.random = random

# Runtime complexity: O(2n) == O(n) (for iterating over all the original nodes, each entry processed twice)
# Auxiliary space complexity: O(n) (for holding all the original nodes)
# Subjective level: medium, bit tedious (as usual with linked lists)
# Solved on: 2026-02-18
class Solution:
    nodes = {}
    def copyRandomList(self, head: Node | None) -> Node | None:
        if head is None: # this covers both the end of a list (`next`) and a `random` pointing to a null.
            return None
        if head in self.nodes:
            return self.nodes[head]

        # otherwise create a new Node, store it in static self.nodes and recursively process it:
        node = Node(head.val)
        self.nodes[head] = node
        node.next = self.copyRandomList(head.next)
        node.random = self.copyRandomList(head.random)
        return node
