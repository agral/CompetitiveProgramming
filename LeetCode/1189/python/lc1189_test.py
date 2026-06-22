import unittest

from lc1189 import Solution

class Test_maxNumberOfBalloons(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxNumberOfBalloons("nlaebolko"), 1)
        self.assertEqual(s.maxNumberOfBalloons("loonbalxballpoon"), 2)
        self.assertEqual(s.maxNumberOfBalloons("leetcode"), 0)
        self.assertEqual(s.maxNumberOfBalloons("bbbbaaaaaallllllllllloooooooooooooon"), 1)
        self.assertEqual(s.maxNumberOfBalloons("bbbbaaaaaallllllllllloooooooooooooonnn"), 3)

if __name__ == "__main__":
    unittest.main()
