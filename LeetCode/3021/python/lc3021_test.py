import unittest

from lc3021 import Solution

class Test_flowerGame(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.flowerGame(3, 2), 3)
        self.assertEqual(s.flowerGame(1, 1), 0)

if __name__ == "__main__":
    unittest.main()
