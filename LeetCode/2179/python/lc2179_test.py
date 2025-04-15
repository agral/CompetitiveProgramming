import unittest

from lc2179 import Solution

class Test_goodTriplets(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.goodTriplets([2, 0, 1, 3], [0, 1, 2, 3]), 1)
        self.assertEqual(s.goodTriplets([4, 0, 1, 3, 2], [4, 1, 0, 2, 3]), 4)

if __name__ == "__main__":
    unittest.main()
