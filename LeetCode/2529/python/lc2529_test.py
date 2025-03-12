import unittest

from lc2529 import Solution

class Test_maximumCount(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maximumCount([-2, -1, -1, 1, 2, 3]), 3)
        self.assertEqual(s.maximumCount([-3, -2, -1, 0, 0, 1, 2]), 3)
        self.assertEqual(s.maximumCount([5, 20, 66, 1314]), 4)

if __name__ == "__main__":
    unittest.main()
