import unittest

from lc3105 import Solution

class Test_longestMonotonicSubarray(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.longestMonotonicSubarray([1, 4, 3, 3, 2]), 2)
        self.assertEqual(s.longestMonotonicSubarray([3, 3, 3, 3]), 1)
        self.assertEqual(s.longestMonotonicSubarray([3, 2, 1]), 3)

if __name__ == "__main__":
    unittest.main()
