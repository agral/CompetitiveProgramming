from lc1019 import Solution, ListNode
from typing import List
import unittest

def makeList(values: List[int]) -> ListNode | None:
    if not values:
        return
    head = ListNode(values[0])
    prev = head
    for i in range(1, len(values)):
        curr = ListNode(values[i])
        prev.next = curr
        prev = curr
    return head

def toString(head: ListNode | None) -> str:
    values = []
    curr = head
    while curr != None:
        values.append(str(curr.val))
        curr = curr.next
    return f"[{' -> '.join(values)}]"


class Test_nextLargerNodes(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.nextLargerNodes(makeList([2, 1, 5])), [5, 5, 0])
        self.assertEqual(s.nextLargerNodes(makeList([2, 7, 4, 3, 5])), [7, 0, 5, 5, 0])

if __name__ == "__main__":
    unittest.main()
