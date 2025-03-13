import unittest

from lc3356 import Solution

class Test_minZeroArray(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minZeroArray([2, 0, 2], [[0, 2, 1], [0, 2, 1], [1, 1, 3]]), 2)
        self.assertEqual(s.minZeroArray([4, 3, 2, 1], [[1, 3, 2], [0, 2, 1]]), -1)

if __name__ == "__main__":
    unittest.main()
