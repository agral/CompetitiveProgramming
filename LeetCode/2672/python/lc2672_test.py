import unittest

from lc2672 import Solution

class Test_colorTheArray(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.colorTheArray(4, [[0, 2], [1, 2], [3, 1], [1, 1], [2, 1]]), [0, 1, 1, 0, 2])
        self.assertEqual(s.colorTheArray(1, [[0, 100_000]]), [0])

if __name__ == "__main__":
    unittest.main()
