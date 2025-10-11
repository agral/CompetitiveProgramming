import unittest

from lc3186 import Solution

class Test_maximumTotalDamage(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maximumTotalDamage([1, 1, 3, 4]), 6)
        self.assertEqual(s.maximumTotalDamage([7, 1, 6, 6]), 13)

if __name__ == "__main__":
    unittest.main()
