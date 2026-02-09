import unittest

from lc2073 import Solution

class Test_timeRequiredToBuy(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.timeRequiredToBuy([2, 3, 2], 2), 6)
        self.assertEqual(s.timeRequiredToBuy([5, 1, 1, 1], 0), 8)
        self.assertEqual(s.timeRequiredToBuy([5, 1, 1, 1], 3), 4)

if __name__ == "__main__":
    unittest.main()
