import unittest

from lc3660 import Solution

class Test_maxValue(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxValue([2, 1, 3]), [2, 2, 3])
        self.assertEqual(s.maxValue([2, 3, 1]), [3, 3, 3])

if __name__ == "__main__":
    unittest.main()
