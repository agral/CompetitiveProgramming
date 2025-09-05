import unittest

from lc2749 import Solution

class Test_makeTheIntegerZero(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.makeTheIntegerZero(5, 7), -1)
        self.assertEqual(s.makeTheIntegerZero(3, -2), 3)

if __name__ == "__main__":
    unittest.main()
