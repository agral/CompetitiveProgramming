import unittest

from lc0941 import Solution

class Test_validMountainArray(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.validMountainArray([2, 1]), False)
        self.assertEqual(s.validMountainArray([3, 5, 5]), False)
        self.assertEqual(s.validMountainArray([0, 3, 2, 1]), True)
        self.assertEqual(s.validMountainArray([0, 1, 2, 3, 4, 5, 6, 7, 8, 9]), False)

if __name__ == "__main__":
    unittest.main()
