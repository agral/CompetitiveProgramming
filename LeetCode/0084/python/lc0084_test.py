import unittest

from lc0084 import Solution

class Test_largestRectangleArea(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.largestRectangleArea([2, 1, 5, 6, 2, 3]), 10)
        self.assertEqual(s.largestRectangleArea([2, 4]), 4)

if __name__ == "__main__":
    unittest.main()
