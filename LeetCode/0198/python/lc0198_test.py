import unittest

from lc0198 import Solution

class Test_rob(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.rob([1, 2, 3, 1]), 4)
        self.assertEqual(s.rob([2, 7, 9, 3, 1]), 12)
        self.assertEqual(s.rob([100, 1, 1, 100]), 200)

if __name__ == "__main__":
    unittest.main()
