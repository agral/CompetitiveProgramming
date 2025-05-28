import unittest

from lc3372 import Solution

class Test_maxTargetNodes(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxTargetNodes(
            [[0, 1], [0, 2], [2, 3], [2, 4]],
            [[0, 1], [0, 2], [0, 3], [2, 7], [1, 4], [4, 5], [4, 6]],
            2
        ), [9, 7, 9, 8, 8])

        self.assertEqual(s.maxTargetNodes(
            [[0, 1], [0, 2], [0, 3], [0, 4]],
            [[0, 1], [1, 2], [2, 3]],
            1
        ), [6, 3, 3, 3, 3])

if __name__ == "__main__":
    unittest.main()
