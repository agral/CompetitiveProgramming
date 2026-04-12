import unittest

from lc1320 import Solution

class Test_minimumDistance(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minimumDistance("CAKE"), 3)
        self.assertEqual(s.minimumDistance("HAPPY"), 6)
        self.assertEqual(s.minimumDistance("ABABZ"), 3)

if __name__ == "__main__":
    unittest.main()
