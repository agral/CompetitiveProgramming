import unittest

from lc1475 import Solution

class Test_finalPrices(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.finalPrices([8, 4, 6, 2, 3]), [4, 2, 4, 2, 3])
        self.assertEqual(s.finalPrices([1, 2, 3, 4, 5]), [1, 2, 3, 4, 5])
        self.assertEqual(s.finalPrices([10, 1, 1, 6]), [9, 0, 1, 6])


if __name__ == "__main__":
    unittest.main()
