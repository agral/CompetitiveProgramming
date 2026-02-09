import unittest

from lc0066 import Solution

class Test_plusOne(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.plusOne([1, 2, 3]), [1, 2, 4])
        self.assertEqual(s.plusOne([4, 3, 2, 1]), [4, 3, 2, 2])
        self.assertEqual(s.plusOne([9]), [1, 0])
        self.assertEqual(s.plusOne([0]), [1])
        self.assertEqual(s.plusOne([1, 9]), [2, 0])
        self.assertEqual(s.plusOne([8, 9]), [9, 0])
        self.assertEqual(s.plusOne([9, 9]), [1, 0, 0])

if __name__ == "__main__":
    unittest.main()
