import unittest

from lc1792 import Solution

class Test_maxAverageRatio(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertAlmostEqual(s.maxAverageRatio([[1, 2], [3, 5], [2, 2]], 2), 0.78333, places=5)
        self.assertAlmostEqual(s.maxAverageRatio([[2, 4], [3, 9], [4, 5], [2, 10]], 4), 0.53485, places=5)

if __name__ == "__main__":
    unittest.main()
