import unittest

from lc2654 import Solution

class Test_minOperations(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.minOperations([2, 6, 3, 4]), 4)
        self.assertEqual(s.minOperations([2, 10, 6, 14]), -1)
        self.assertEqual(s.minOperations([6, 10, 15]), 4)

if __name__ == "__main__":
    unittest.main()
