from lc0206 import Solution, ListNode
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


class Test_reverseList(unittest.TestCase):
    def test(self):
        s = Solution()

        self.assertEqual(toString(s.reverseList(makeList([1, 2, 3, 4, 5]))),
                         toString(makeList([5, 4, 3, 2, 1])))
        self.assertEqual(toString(s.reverseList(makeList([1, 2]))),
                         toString(makeList([2, 1])))
        self.assertEqual(toString(s.reverseList(makeList([]))),
                         toString(makeList([])))

if __name__ == "__main__":
    unittest.main()
