import unittest

from lc2176 import Solution

class Test_countPairs(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.countPairs([3, 1, 2, 2, 2, 1, 3], 2), 4)
        self.assertEqual(s.countPairs([1, 2, 3, 4], 1), 0)

if __name__ == "__main__":
    unittest.main()
