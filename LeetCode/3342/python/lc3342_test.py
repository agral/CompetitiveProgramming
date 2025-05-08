import unittest

from lc3342 import Solution

class Test_minTimeToReach(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minTimeToReach([[0, 4], [4, 4]]), 7)
        self.assertEqual(s.minTimeToReach([[0, 0, 0, 0], [0, 0, 0, 0]]), 6)
        self.assertEqual(s.minTimeToReach([[0, 1], [1, 2]]), 4)

if __name__ == "__main__":
    unittest.main()
