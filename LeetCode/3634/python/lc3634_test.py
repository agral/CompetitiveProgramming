import unittest

from lc3634 import Solution

class Test_minRemoval(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minRemoval([2, 1, 5], 2), 1)    # Remove 5, get [2, 1], 1*2 <= 2.
        self.assertEqual(s.minRemoval([1, 6, 2, 9], 3), 2) # Remove 6,9, get[1, 2], 1*3 <= 2.
        self.assertEqual(s.minRemoval([4, 6], 2), 0)       # Get [4, 6]; 4*2 <= 6
        self.assertEqual(s.minRemoval([42], 1), 0)         # Array of length 1 -> no removals needed.
        self.assertEqual(s.minRemoval([466,306,76,17,60,246,341,284], 2), 3) # WA #01

if __name__ == "__main__":
    unittest.main()
