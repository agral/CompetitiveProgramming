import unittest

from lc2078 import Solution

class Test_maxDistance(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxDistance([1, 1, 1, 6, 1, 1, 1]), 3)
        self.assertEqual(s.maxDistance([1, 8, 3, 8, 3]), 4)
        self.assertEqual(s.maxDistance([0, 1]), 1)
        self.assertEqual(s.maxDistance([4,4,4,11,4,4,11,4,4,4,4,4]), 8)

if __name__ == "__main__":
    unittest.main()
