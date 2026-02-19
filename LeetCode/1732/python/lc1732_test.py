import unittest

from lc1732 import Solution

class Test_largestAltitude(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.largestAltitude([-5, 1, 5, 0, -7]), 1)
        self.assertEqual(s.largestAltitude([-4, -3, -2, -1, 4, 3, 2]), 0)

if __name__ == "__main__":
    unittest.main()
