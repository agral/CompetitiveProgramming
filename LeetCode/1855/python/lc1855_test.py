import unittest

from lc1855 import Solution

class Test_maxDistance(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxDistance([55, 30, 5, 4, 2], [100, 20, 10, 10, 5]), 2)
        self.assertEqual(s.maxDistance([2, 2, 2], [10, 10, 1]), 1)
        self.assertEqual(s.maxDistance([30, 29, 19, 5], [25, 25, 25, 25, 25]), 2)

if __name__ == "__main__":
    unittest.main()
