import unittest

from lc2141 import Solution

class Test_maxRunTime(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxRunTime(2, [3, 3, 3]), 4)
        self.assertEqual(s.maxRunTime(2, [1, 1, 1, 1]), 2)

if __name__ == "__main__":
    unittest.main()
