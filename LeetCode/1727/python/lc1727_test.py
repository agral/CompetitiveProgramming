import unittest

from lc1727 import Solution

class Test_largestSubmatrix(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.largestSubmatrix([
                [0, 0, 1],
                [1, 1, 1],
                [1, 0, 1],
        ]), 4)

        self.assertEqual(s.largestSubmatrix([
                [1, 0, 1, 0, 1],
        ]), 3)

        self.assertEqual(s.largestSubmatrix([
                [1, 1, 0],
                [1, 0, 1],
        ]), 2)

if __name__ == "__main__":
    unittest.main()
