import unittest

from lc0852 import Solution

class Test_peakIndexInMountainArray(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.peakIndexInMountainArray([0, 1, 0]), 1)
        self.assertEqual(s.peakIndexInMountainArray([0, 2, 1, 0]), 1)
        self.assertEqual(s.peakIndexInMountainArray([0, 10, 5, 2]), 1)

if __name__ == "__main__":
    unittest.main()
