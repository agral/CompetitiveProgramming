import unittest

from lc0200 import Solution

class Test_numIslands(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.numIslands([
            ["1","1","1","1","0"],
            ["1","1","0","1","0"],
            ["1","1","0","0","0"],
            ["0","0","0","0","0"]
        ]), 1)

        self.assertEqual(s.numIslands([
            ["1","1","0","0","0"],
            ["1","1","0","0","0"],
            ["0","0","1","0","0"],
            ["0","0","0","1","1"]
        ]), 3)

if __name__ == "__main__":
    unittest.main()
