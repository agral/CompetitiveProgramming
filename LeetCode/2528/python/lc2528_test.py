import unittest

from lc2528 import Solution

class Test_maxPower(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxPower([1, 2, 4, 5, 0], 1, 2), 5)
        self.assertEqual(s.maxPower([4, 4, 4, 4], 0, 3), 4)

if __name__ == "__main__":
    unittest.main()
