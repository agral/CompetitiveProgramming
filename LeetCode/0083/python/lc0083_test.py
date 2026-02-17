import unittest

from lc0083 import *

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

class Test_deleteDuplicates(unittest.TestCase):
    def test(self):
        s = Solution()

        input = makeList([1, 1, 2])
        actual = s.deleteDuplicates(input)
        self.assertEqual(toString(actual), toString(makeList([1, 2])))

        input = makeList([1, 1, 2, 3, 3])
        actual = s.deleteDuplicates(input)
        self.assertEqual(toString(actual), toString(makeList([1, 2, 3])))

        input = makeList([])
        actual = s.deleteDuplicates(input)
        self.assertEqual(toString(actual), toString(makeList([])))

        input = makeList([1, 1, 1])
        actual = s.deleteDuplicates(input)
        self.assertEqual(toString(actual), toString(makeList([1])))

if __name__ == "__main__":
    unittest.main()
