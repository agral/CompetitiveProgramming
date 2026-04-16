import unittest

from lc0121 import Solution

class Test_maxProfit(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxProfit([7, 1, 5, 3, 6, 4]), 5)
        self.assertEqual(s.maxProfit([7, 6, 4, 3, 1]), 0)

if __name__ == "__main__":
    unittest.main()
