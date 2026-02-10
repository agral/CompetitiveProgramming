import unittest

from lc1046 import Solution

class Test_lastStoneWeight(unittest.TestCase):
    def test(self):
        s = Solution()
        # Example 1: [8, 7, 4, 2, 1, 1] -> [4, 2, 1, 1, 1], [2, 2, 1, 1, 1] -> [1, 1, 1] -> [1].
        self.assertEqual(s.lastStoneWeight([2, 7, 4, 1, 8, 1]), 1)

        # Example 2: [1]
        self.assertEqual(s.lastStoneWeight([1]), 1)



if __name__ == "__main__":
    unittest.main()
