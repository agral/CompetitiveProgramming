import unittest

from lc0328 import Solution, ListNode
from typing import List

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

class Test_oddEvenList(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(toString(s.oddEvenList(makeList([1, 2, 3, 4, 5]))),
                         toString(makeList([1, 3, 5, 2, 4])))
        self.assertEqual(toString(s.oddEvenList(makeList([2, 1, 3, 5, 6, 4, 7]))),
                         toString(makeList([2, 3, 6, 7, 1, 5, 4])))

if __name__ == "__main__":
    unittest.main()
