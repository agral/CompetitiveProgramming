import unittest

from typing import List
from lc0147 import Solution, ListNode

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

class Test_insertionSortList(unittest.TestCase):
    def test(self):
        s = Solution()

        example1 = makeList([4, 2, 1, 3])
        expected1 = makeList([1, 2, 3, 4])
        self.assertEqual(toString(s.insertionSortList(example1)), toString(expected1))
        example2 = makeList([-1, 5, 3, 4, 0])
        expected2 = makeList([-1, 0, 3, 4, 5])
        self.assertEqual(toString(s.insertionSortList(example2)), toString(expected2))

if __name__ == "__main__":
    unittest.main()
