import unittest

from lc3373 import Solution

class Test_maxTargetNodes(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxTargetNodes(
            [[0, 1], [0, 2], [2, 3], [2, 4]],
            [[0, 1], [0, 2], [0, 3], [2, 7], [1, 4], [4, 5], [4, 6]]
        ), [8, 7, 7, 8, 8])

        self.assertEqual(s.maxTargetNodes(
            [[0, 1], [0, 2], [0, 3], [0, 4]],
            [[0, 1], [1, 2], [2, 3]],
        ), [3, 6, 6, 6, 6])

if __name__ == "__main__":
    unittest.main()
