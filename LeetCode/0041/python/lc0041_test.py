import unittest

from lc0041 import Solution

class Test_firstMissingPositive(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.firstMissingPositive([2, 1, 0]), 3)
        self.assertEqual(s.firstMissingPositive([0, 1, 2]), 3)
        self.assertEqual(s.firstMissingPositive([3, 4, -1, 1]), 2)
        self.assertEqual(s.firstMissingPositive([7, 8, 9, 11, 12]), 1)

if __name__ == "__main__":
    unittest.main()
