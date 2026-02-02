import unittest

from lc3013 import Solution

class Test_minimumCost(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minimumCost([1, 3, 2, 6, 4, 2], 3, 3), 5)
        self.assertEqual(s.minimumCost([10, 1, 2, 2, 2, 1], 4, 3), 15)
        self.assertEqual(s.minimumCost([10, 8, 18, 9], 3, 1), 36)

if __name__ == "__main__":
    unittest.main()
