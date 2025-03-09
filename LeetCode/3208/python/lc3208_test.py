import unittest

from lc3208 import Solution

class Test_numberOfAlternatingGroups(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.numberOfAlternatingGroups([0, 1, 0, 1, 0], 3), 3)
        self.assertEqual(s.numberOfAlternatingGroups([0, 1, 0, 0, 1, 0, 1], 6), 2)
        self.assertEqual(s.numberOfAlternatingGroups([1, 1, 0, 1], 4), 0)

if __name__ == "__main__":
    unittest.main()
