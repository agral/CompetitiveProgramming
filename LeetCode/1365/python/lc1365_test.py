import unittest

from lc1365 import Solution

class Test_smallerNumbersThanCurrent(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.smallerNumbersThanCurrent([8, 1, 2, 2, 3]), [4, 0, 1, 1, 3])
        self.assertEqual(s.smallerNumbersThanCurrent([6, 5, 4, 8]), [2, 1, 0, 3])
        self.assertEqual(s.smallerNumbersThanCurrent([7, 7, 7, 7]), [0, 0, 0, 0])

if __name__ == "__main__":
    unittest.main()
