import unittest

from lc0739 import Solution

class Test_dailyTemperatures(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.dailyTemperatures(
            [73, 74, 75, 71, 69, 72, 76, 73]),
            [ 1,  1,  4,  2,  1,  1,  0,  0])
        self.assertEqual(s.dailyTemperatures(
            [30, 40, 50, 60]),
            [ 1,  1,  1,  0])
        self.assertEqual(s.dailyTemperatures(
            [30, 60, 90]),
            [ 1,  1,  0]),

if __name__ == "__main__":
    unittest.main()
