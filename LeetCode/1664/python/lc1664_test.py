import unittest

from lc1664 import Solution

class Test_waysToMakeFair(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.waysToMakeFair([6, 1, 7, 4, 1]), 0)
        self.assertEqual(s.waysToMakeFair([2, 1, 6, 4]), 1)
        self.assertEqual(s.waysToMakeFair([1, 1, 1]), 3)
        self.assertEqual(s.waysToMakeFair([1, 2, 3]), 0)


if __name__ == "__main__":
    unittest.main()
