import unittest

from lc3027 import Solution

class Test_numberOfPairs(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.numberOfPairs([[1, 1], [2, 2], [3, 3]]), 0)
        self.assertEqual(s.numberOfPairs([[6, 2], [4, 4], [2, 6]]), 2)
        self.assertEqual(s.numberOfPairs([[3, 1], [1, 3], [1, 1]]), 2)

if __name__ == "__main__":
    unittest.main()
