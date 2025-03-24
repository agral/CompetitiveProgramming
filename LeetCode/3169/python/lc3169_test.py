import unittest

from lc3169 import Solution

class Test_countDays(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.countDays(10, [[5, 7], [1, 3], [9, 10]]), 2)
        self.assertEqual(s.countDays(5, [[2, 4], [1, 3]]), 1)
        self.assertEqual(s.countDays(6, [[1, 6]]), 0)

if __name__ == "__main__":
    unittest.main()
