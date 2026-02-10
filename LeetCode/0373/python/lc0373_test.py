import unittest

from lc0373 import Solution

class Test_kSmallestPairs(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.kSmallestPairs(
            [1, 7, 11], [2, 4, 6], 3),
            [[1, 2], [1, 4], [1, 6]]
        )
        self.assertEqual(s.kSmallestPairs(
            [1, 1, 2], [1, 1, 3], 2),
            [[1, 1], [1, 1]]
        )

if __name__ == "__main__":
    unittest.main()
