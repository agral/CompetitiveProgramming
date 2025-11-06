import unittest

from lc3607 import Solution

class Test_processQueries(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.processQueries(
            5,
            [[1,2],[2,3],[3,4],[4,5]],
            [[1,3],[2,1],[1,1],[2,2],[1,2]]
        ), [3, 2, 3])

        self.assertEqual(s.processQueries(
            3,
            [],
            [[1,1],[2,1],[1,1]]
        ), [1, -1])

if __name__ == "__main__":
    unittest.main()
