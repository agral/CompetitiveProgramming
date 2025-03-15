import unittest

from lc2560 import Solution

class Test_minCapability(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minCapability([2, 3, 5, 9], 2), 5)
        self.assertEqual(s.minCapability([2, 7, 9, 3, 1], 2), 2)

if __name__ == "__main__":
    unittest.main()
