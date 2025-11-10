import unittest

from lc3542 import Solution

class Test_minOperations(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minOperations([0, 2]), 1)
        self.assertEqual(s.minOperations([3, 1, 2, 1]), 3)
        self.assertEqual(s.minOperations([1, 2, 1, 2, 1, 2]), 4)

if __name__ == "__main__":
    unittest.main()
