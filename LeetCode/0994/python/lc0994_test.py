import unittest

from lc0994 import Solution

class Test_orangesRotting(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.orangesRotting([
            [2, 1, 1],
            [1, 1, 0],
            [0, 1, 1]]), 4)
        self.assertEqual(s.orangesRotting([
            [2, 1, 1],
            [0, 1, 1],
            [1, 0, 1]]), -1) # the one in the bottom left corner never rots.
        self.assertEqual(s.orangesRotting([
            [0, 2]]), 0) # no fresh oranges to begin with; so the answer is 0.
        self.assertEqual(s.orangesRotting([
            [2, 1, 1],
            [1, 1, 1],
            [1, 1, 1]]), 4)

if __name__ == "__main__":
    unittest.main()
