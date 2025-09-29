import unittest

from lc1039 import Solution

class Test_minScoreTriangulation(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minScoreTriangulation([1, 2, 3]), 6)
        self.assertEqual(s.minScoreTriangulation([3, 7, 4, 5]), 144)
        self.assertEqual(s.minScoreTriangulation([1, 3, 1, 4, 1, 5]), 13)

if __name__ == "__main__":
    unittest.main()
