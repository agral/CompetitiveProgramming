import unittest

from lc2542 import Solution

class Test_maxScore(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxScore([1, 3, 3, 2], [2, 1, 3, 4], 3), 12)
        self.assertEqual(s.maxScore([4, 2, 3, 1, 1], [7, 5, 10, 9, 6], 1), 30)

if __name__ == "__main__":
    unittest.main()
