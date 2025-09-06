import unittest

from lc3495 import Solution

class Test_minOperations(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minOperations([[1, 2], [2, 4]]), 3)
        self.assertEqual(s.minOperations([[2, 6]]), 4)

if __name__ == "__main__":
    unittest.main()
