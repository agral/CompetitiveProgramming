import unittest

from lc0912 import Solution

class Test_sortArray(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.sortArray([5, 2, 3, 1]), [1, 2, 3, 5])
        self.assertEqual(s.sortArray([5, 1, 1, 2, 0, 0]), [0, 0, 1, 1, 2, 5])
        self.assertEqual(s.sortArray([-2, 3, -5]), [-5, -2, 3])

if __name__ == "__main__":
    unittest.main()
