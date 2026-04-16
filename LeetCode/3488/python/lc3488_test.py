import unittest

from lc3488 import Solution

class Test_solveQueries(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.solveQueries([1, 3, 1, 4, 1, 3, 2], [0, 3, 5]), [2, -1, 3])
        self.assertEqual(s.solveQueries([1, 2, 3, 4], [0, 1, 2, 3]), [-1, -1, -1, -1])

if __name__ == "__main__":
    unittest.main()
