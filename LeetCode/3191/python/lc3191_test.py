import unittest

from lc3191 import Solution

class Test_minOperations(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minOperations([0, 1, 1, 1, 0, 0]), 3)
        self.assertEqual(s.minOperations([0, 1, 1, 1]), -1)

if __name__ == "__main__":
    unittest.main()
