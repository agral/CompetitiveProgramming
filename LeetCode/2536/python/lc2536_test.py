import unittest

from lc2536 import Solution

class Test_rangeAddQueries(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.rangeAddQueries(3, [[1, 1, 2, 2], [0, 0, 1, 1]]),
                         [[1, 1, 0], [1, 2, 1], [0, 1, 1]])
        self.assertEqual(s.rangeAddQueries(2, [[0, 0, 1, 1]]), [[1, 1], [1, 1]])

if __name__ == "__main__":
    unittest.main()
