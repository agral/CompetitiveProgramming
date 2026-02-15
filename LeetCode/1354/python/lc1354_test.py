import unittest

from lc1354 import Solution

class Test_isPossible(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertTrue(s.isPossible([9, 3, 5])) # Example 1
        self.assertFalse(s.isPossible([1, 1, 1, 2])) # Example 2
        self.assertTrue(s.isPossible([8, 5])) # Example 3
        self.assertFalse(s.isPossible([1, 1, 2]))
        self.assertTrue(s.isPossible([1, 1, 3]))
        self.assertFalse(s.isPossible([1, 1, 4]))
        self.assertTrue(s.isPossible([7, 1, 1, 4]))
        self.assertFalse(s.isPossible([8, 1, 1, 4]))

if __name__ == "__main__":
    unittest.main()
