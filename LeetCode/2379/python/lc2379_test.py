import unittest

from lc2379 import Solution

class Test_minimumRecolors(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minimumRecolors("WBBWWBBWBW", 7), 3)
        self.assertEqual(s.minimumRecolors("WBWBBBW", 2), 0)

if __name__ == "__main__":
    unittest.main()
