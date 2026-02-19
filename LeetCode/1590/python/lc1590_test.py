import unittest

from lc1590 import Solution

class Test_minSubarray(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minSubarray([3, 1, 4, 2], 6), 1) # remove [4], [3, 1, 2] remains.
        self.assertEqual(s.minSubarray([6, 3, 5, 2], 9), 2) # remove [5, 2], [6, 3] remains.
        self.assertEqual(s.minSubarray([1, 2, 3], 3), 0)    # remove [], [1, 2, 3] remains.
        self.assertEqual(s.minSubarray([1, 2, 3, 4], 11), -1) # no way to remove a subarray.


if __name__ == "__main__":
    unittest.main()
